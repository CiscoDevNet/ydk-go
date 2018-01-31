// This module contains a collection of YANG definitions
// for Cisco IOS-XR ethernet-cfm package operational data.
// 
// This module contains definitions
// for the following management objects:
//   cfm: CFM operational data
// 
// This YANG module augments the
//   Cisco-IOS-XR-infra-sla-oper
// module with state data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_cfm_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_cfm_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ethernet-cfm-oper cfm}", reflect.TypeOf(Cfm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ethernet-cfm-oper:cfm", reflect.TypeOf(Cfm{}))
}

// CfmPmElrIngressAction represents ELR Ingress action
type CfmPmElrIngressAction string

const (
    // OK
    CfmPmElrIngressAction_elr_ingress_ok CfmPmElrIngressAction = "elr-ingress-ok"

    // Down
    CfmPmElrIngressAction_elr_ingress_down CfmPmElrIngressAction = "elr-ingress-down"

    // STP Blocked
    CfmPmElrIngressAction_elr_ingress_blocked CfmPmElrIngressAction = "elr-ingress-blocked"

    // VID Blocked
    CfmPmElrIngressAction_elr_ingress_vid CfmPmElrIngressAction = "elr-ingress-vid"
)

// CfmPmRelayAction represents LTR relay action
type CfmPmRelayAction string

const (
    // Target Hit
    CfmPmRelayAction_relay_hit CfmPmRelayAction = "relay-hit"

    // Filtering database
    CfmPmRelayAction_relay_fdb CfmPmRelayAction = "relay-fdb"

    // CCM Learning database
    CfmPmRelayAction_relay_mpdb CfmPmRelayAction = "relay-mpdb"
)

// CfmBagSmanFmt represents Short MA Name format
type CfmBagSmanFmt string

const (
    // Short MA Name is a 12-bit VLAN-ID
    CfmBagSmanFmt_sman_vlan_id CfmBagSmanFmt = "sman-vlan-id"

    // Short MA Name is a character string
    CfmBagSmanFmt_sman_string CfmBagSmanFmt = "sman-string"

    // Short MA Name is a 16-bit unsigned integer
    CfmBagSmanFmt_sman_uint16 CfmBagSmanFmt = "sman-uint16"

    // Short MA Name is a global VPN identifier
    CfmBagSmanFmt_sman_vpn_id CfmBagSmanFmt = "sman-vpn-id"

    // Short MA Name uses the ICC-based format
    CfmBagSmanFmt_sman_icc CfmBagSmanFmt = "sman-icc"

    // Unknown Short MA Name format
    CfmBagSmanFmt_sman_unknown CfmBagSmanFmt = "sman-unknown"
)

// CfmPmMepDefect represents Defects that can be reported by a MEP
type CfmPmMepDefect string

const (
    // No defect reported
    CfmPmMepDefect_defect_none CfmPmMepDefect = "defect-none"

    // Some Peer MEP's CCM has the RDI bit set
    CfmPmMepDefect_defect_rdi_ccm CfmPmMepDefect = "defect-rdi-ccm"

    // A Peer MEP port or interface status error has
    // been reported
    CfmPmMepDefect_defect_ma_cstatus CfmPmMepDefect = "defect-ma-cstatus"

    // Not receiving valid CCMs from at least one Peer
    // MEP
    CfmPmMepDefect_defect_remote_ccm CfmPmMepDefect = "defect-remote-ccm"

    // Currently receiving invalid CCMs from at least
    // one Peer MEP
    CfmPmMepDefect_defect_error_ccm CfmPmMepDefect = "defect-error-ccm"

    // Currently receiving CCMs from an incorrect
    // service (MA)
    CfmPmMepDefect_defect_cross_connect_ccm CfmPmMepDefect = "defect-cross-connect-ccm"
)

// CfmPmElrEgressAction represents ELR Egress action
type CfmPmElrEgressAction string

const (
    // OK
    CfmPmElrEgressAction_elr_egress_ok CfmPmElrEgressAction = "elr-egress-ok"

    // Down
    CfmPmElrEgressAction_elr_egress_down CfmPmElrEgressAction = "elr-egress-down"

    // STP Blocked
    CfmPmElrEgressAction_elr_egress_blocked CfmPmElrEgressAction = "elr-egress-blocked"

    // VID Blocked
    CfmPmElrEgressAction_elr_egress_vid CfmPmElrEgressAction = "elr-egress-vid"

    // MAC Pruned
    CfmPmElrEgressAction_elr_egress_mac CfmPmElrEgressAction = "elr-egress-mac"
)

// CfmPmIngressAction represents Ingress action
type CfmPmIngressAction string

const (
    // OK
    CfmPmIngressAction_ingress_ok CfmPmIngressAction = "ingress-ok"

    // Down
    CfmPmIngressAction_ingress_down CfmPmIngressAction = "ingress-down"

    // STP Blocked
    CfmPmIngressAction_ingress_blocked CfmPmIngressAction = "ingress-blocked"

    // VID Blocked
    CfmPmIngressAction_ingress_vid CfmPmIngressAction = "ingress-vid"
)

// CfmBagCcmInterval represents CFM CCM intervals
type CfmBagCcmInterval string

const (
    // Invalid CCM interval
    CfmBagCcmInterval_interval_none CfmBagCcmInterval = "interval-none"

    // Interval of 3.3ms
    CfmBagCcmInterval_interval3_3ms CfmBagCcmInterval = "interval3-3ms"

    // Interval of 10ms
    CfmBagCcmInterval_interval10ms CfmBagCcmInterval = "interval10ms"

    // Interval of 100ms
    CfmBagCcmInterval_interval100ms CfmBagCcmInterval = "interval100ms"

    // Interval of 1s
    CfmBagCcmInterval_interval1s CfmBagCcmInterval = "interval1s"

    // Interval of 10s
    CfmBagCcmInterval_interval10s CfmBagCcmInterval = "interval10s"

    // Interval of 1 min
    CfmBagCcmInterval_interval1m CfmBagCcmInterval = "interval1m"

    // Interval of 10 mins
    CfmBagCcmInterval_interval10m CfmBagCcmInterval = "interval10m"
)

// CfmPmChassisIdFmt represents Chassis ID type
type CfmPmChassisIdFmt string

const (
    // Chassis ID is a component name
    CfmPmChassisIdFmt_chassis_id_chassis_component CfmPmChassisIdFmt = "chassis-id-chassis-component"

    // Chassis ID is an interface alias
    CfmPmChassisIdFmt_chassis_id_interface_alias CfmPmChassisIdFmt = "chassis-id-interface-alias"

    // Chassis ID is a port component name
    CfmPmChassisIdFmt_chassis_id_port_component CfmPmChassisIdFmt = "chassis-id-port-component"

    // Chassis ID is a MAC address
    CfmPmChassisIdFmt_chassis_id_mac_address CfmPmChassisIdFmt = "chassis-id-mac-address"

    // Chassis ID is a network address
    CfmPmChassisIdFmt_chassis_id_network_address CfmPmChassisIdFmt = "chassis-id-network-address"

    // Chassis ID is an interface name
    CfmPmChassisIdFmt_chassis_id_interface_name CfmPmChassisIdFmt = "chassis-id-interface-name"

    // Chassis ID is a local name
    CfmPmChassisIdFmt_chassis_id_local CfmPmChassisIdFmt = "chassis-id-local"

    // Unknown Chassis ID type
    CfmPmChassisIdFmt_chassis_id_unknown_type CfmPmChassisIdFmt = "chassis-id-unknown-type"
)

// SlaOperOperation represents Type of SLA operation
type SlaOperOperation string

const (
    // Configured SLA operation
    SlaOperOperation_operation_type_configured SlaOperOperation = "operation-type-configured"

    // On-demand SLA operation
    SlaOperOperation_operation_type_ondemand SlaOperOperation = "operation-type-ondemand"
)

// CfmPmLastHopFmt represents Last hop identifier format
type CfmPmLastHopFmt string

const (
    // No last hop identifier
    CfmPmLastHopFmt_last_hop_none CfmPmLastHopFmt = "last-hop-none"

    // Last hop identifier is a hostname
    CfmPmLastHopFmt_last_hop_host_name CfmPmLastHopFmt = "last-hop-host-name"

    // Last hop identifier is an egress ID
    CfmPmLastHopFmt_last_hop_egress_id CfmPmLastHopFmt = "last-hop-egress-id"
)

// CfmPmIdFmt represents ID format
type CfmPmIdFmt string

const (
    // ID format is a string
    CfmPmIdFmt_id_format_is_string CfmPmIdFmt = "id-format-is-string"

    // ID format is a MAC address
    CfmPmIdFmt_id_format_is_mac_address CfmPmIdFmt = "id-format-is-mac-address"

    // ID format is raw hex
    CfmPmIdFmt_id_format_is_raw_hex CfmPmIdFmt = "id-format-is-raw-hex"
)

// CfmPmRmepState represents State of the Peer MEP state machine
type CfmPmRmepState string

const (
    // Momentary state during reset
    CfmPmRmepState_peer_mep_idle CfmPmRmepState = "peer-mep-idle"

    // Loss timer not expired since reset, but no
    // valid CCM received
    CfmPmRmepState_peer_mep_start CfmPmRmepState = "peer-mep-start"

    // Loss timer has expired
    CfmPmRmepState_peer_mep_failed CfmPmRmepState = "peer-mep-failed"

    // Loss timer has not expired since last valid CCM
    CfmPmRmepState_peer_mep_ok CfmPmRmepState = "peer-mep-ok"
)

// CfmBagCcmOffload represents Offload status of CCM processing
type CfmBagCcmOffload string

const (
    // CCM processing has not been offloaded
    CfmBagCcmOffload_offload_none CfmBagCcmOffload = "offload-none"

    // CCM processing has been offloaded to software
    CfmBagCcmOffload_offload_software CfmBagCcmOffload = "offload-software"

    // CCM processing has been offloaded to hardware
    CfmBagCcmOffload_offload_hardware CfmBagCcmOffload = "offload-hardware"
)

// CfmPmAisReceive represents signal, directly or via AIS or LCK messages.
type CfmPmAisReceive string

const (
    // No signal received
    CfmPmAisReceive_receive_none CfmPmAisReceive = "receive-none"

    // Receiving AIS messages
    CfmPmAisReceive_receive_ais CfmPmAisReceive = "receive-ais"

    // Receiving LCK messages
    CfmPmAisReceive_receive_lck CfmPmAisReceive = "receive-lck"

    // Receiving AIS directly from another MEP on the
    // same interface
    CfmPmAisReceive_receive_direct CfmPmAisReceive = "receive-direct"
)

// CfmMaMpVariety represents CFM MA Maintenance Point varieties
type CfmMaMpVariety string

const (
    // MIP
    CfmMaMpVariety_mip CfmMaMpVariety = "mip"

    // Up MEP
    CfmMaMpVariety_up_mep CfmMaMpVariety = "up-mep"

    // Down MEP
    CfmMaMpVariety_downmep CfmMaMpVariety = "downmep"

    // Unknown MEP
    CfmMaMpVariety_unknown_mep CfmMaMpVariety = "unknown-mep"
)

// CfmPmPktAction represents Action taken for received packet
type CfmPmPktAction string

const (
    // Packet processed successfully
    CfmPmPktAction_packet_processed CfmPmPktAction = "packet-processed"

    // Packet forwarded
    CfmPmPktAction_packet_forwarded CfmPmPktAction = "packet-forwarded"

    // Packet dropped at a MEP due to unknown opcode
    CfmPmPktAction_unknown_opcode CfmPmPktAction = "unknown-opcode"

    // Packet dropped due to level/opcode filtering at
    // a MEP
    CfmPmPktAction_filter_level CfmPmPktAction = "filter-level"

    // Packet dropped because interface is STP blocked
    CfmPmPktAction_filter_blocked CfmPmPktAction = "filter-blocked"

    // Packet dropped due to local destination MAC
    CfmPmPktAction_filter_local_mac CfmPmPktAction = "filter-local-mac"

    // CCM too short or too long
    CfmPmPktAction_malformed_ccm_size CfmPmPktAction = "malformed-ccm-size"

    // Invalid MEP-ID
    CfmPmPktAction_malformed_ccm_mep_id CfmPmPktAction = "malformed-ccm-mep-id"

    // Packet too short
    CfmPmPktAction_malformed_too_short CfmPmPktAction = "malformed-too-short"

    // Destination MAC address does not match
    // interface
    CfmPmPktAction_malformed_destination_mac_unicast CfmPmPktAction = "malformed-destination-mac-unicast"

    // Invalid multicast destination MAC address
    CfmPmPktAction_malformed_destination_mac_multicast CfmPmPktAction = "malformed-destination-mac-multicast"

    // TLV offset too short or beyond the end of the
    // packet
    CfmPmPktAction_malformed_tlv_offset CfmPmPktAction = "malformed-tlv-offset"

    // Invalid source MAC address for LBM
    CfmPmPktAction_malformed_lbm_source_mac CfmPmPktAction = "malformed-lbm-source-mac"

    // Unknown LTR relay action
    CfmPmPktAction_malformed_ltr_relay_action CfmPmPktAction = "malformed-ltr-relay-action"

    // LTR has neither reply-ingress or reply-egress
    CfmPmPktAction_malformed_ltr_reply_tlv CfmPmPktAction = "malformed-ltr-reply-tlv"

    // Invalid Linktrace Message origin MAC address
    CfmPmPktAction_malformed_lt_origin CfmPmPktAction = "malformed-lt-origin"

    // Invalid LTM target MAC address
    CfmPmPktAction_malformed_ltm_target CfmPmPktAction = "malformed-ltm-target"

    // Invalid source MAC address
    CfmPmPktAction_malformed_source_mac CfmPmPktAction = "malformed-source-mac"

    // Packet too short for CFM header
    CfmPmPktAction_malformed_header_too_short CfmPmPktAction = "malformed-header-too-short"

    // TLV header extends beyond the end of the packet
    CfmPmPktAction_malformed_tlv_header_overrun CfmPmPktAction = "malformed-tlv-header-overrun"

    // TLV extends beyond the end of the packet
    CfmPmPktAction_malformed_tlv_overrun CfmPmPktAction = "malformed-tlv-overrun"

    // Multiple Sender-ID TLVs found
    CfmPmPktAction_malformed_duplicate_sender_id CfmPmPktAction = "malformed-duplicate-sender-id"

    // Multiple Port-status TLVs found
    CfmPmPktAction_malformed_duplicate_port_status CfmPmPktAction = "malformed-duplicate-port-status"

    // Multiple Interface-state TLVs found
    CfmPmPktAction_malformed_duplicate_interface_status CfmPmPktAction = "malformed-duplicate-interface-status"

    // Invalid TLV for this type of packet found
    CfmPmPktAction_malformed_wrong_tlv CfmPmPktAction = "malformed-wrong-tlv"

    // Multiple Data TLVs found
    CfmPmPktAction_malformed_duplicate_data CfmPmPktAction = "malformed-duplicate-data"

    // Multiple LTR-Egress-ID TLVs found
    CfmPmPktAction_malformed_duplicate_ltr_egress_id CfmPmPktAction = "malformed-duplicate-ltr-egress-id"

    // Multiple Reply-ingress TLVs found
    CfmPmPktAction_malformed_duplicate_reply_ingress CfmPmPktAction = "malformed-duplicate-reply-ingress"

    // Multiple Reply-egress TLVs found
    CfmPmPktAction_malformed_duplicate_reply_egress CfmPmPktAction = "malformed-duplicate-reply-egress"

    // Multiple LTM-Egress-ID TLVs found
    CfmPmPktAction_malformed_duplicate_ltm_egress_id CfmPmPktAction = "malformed-duplicate-ltm-egress-id"

    // Sender-ID TLV is too short
    CfmPmPktAction_malformed_sender_id_size CfmPmPktAction = "malformed-sender-id-size"

    // Sender-ID TLV is too short to contain the
    // Chassis ID
    CfmPmPktAction_malformed_chassis_id_size CfmPmPktAction = "malformed-chassis-id-size"

    // Sender-ID TLV is too short to contain the
    // management address domain
    CfmPmPktAction_malformed_mgmt_address_domain_size CfmPmPktAction = "malformed-mgmt-address-domain-size"

    // Sender-ID TLV is too short to contain the
    // management address
    CfmPmPktAction_malformed_mgmt_address_size CfmPmPktAction = "malformed-mgmt-address-size"

    // Port-status TLV is too short
    CfmPmPktAction_malformed_port_status_size CfmPmPktAction = "malformed-port-status-size"

    // Invalid Port status value
    CfmPmPktAction_malformed_port_status CfmPmPktAction = "malformed-port-status"

    // Interface-status TLV is too short
    CfmPmPktAction_malformed_interface_status_size CfmPmPktAction = "malformed-interface-status-size"

    // Invalid Interface status value
    CfmPmPktAction_malformed_interface_status CfmPmPktAction = "malformed-interface-status"

    // Organization-specific TLV is too short
    CfmPmPktAction_malformed_organization_specific_tlv_size CfmPmPktAction = "malformed-organization-specific-tlv-size"

    // Multiple MEP-name TLVs found
    CfmPmPktAction_malformed_duplicate_mep_name CfmPmPktAction = "malformed-duplicate-mep-name"

    // Multiple additional-interface-status TLVs found
    CfmPmPktAction_malformed_duplicate_additional_interface_status CfmPmPktAction = "malformed-duplicate-additional-interface-status"

    // LTR-Egress-ID TLV is too short
    CfmPmPktAction_malformed_ltr_egress_id_size CfmPmPktAction = "malformed-ltr-egress-id-size"

    // Reply-ingress TLV is too short
    CfmPmPktAction_malformed_reply_ingress_size CfmPmPktAction = "malformed-reply-ingress-size"

    // Invalid ingress-action value
    CfmPmPktAction_malformed_ingress_action CfmPmPktAction = "malformed-ingress-action"

    // Reply-ingress TLV has invalid MAC address
    CfmPmPktAction_malformed_reply_ingress_mac CfmPmPktAction = "malformed-reply-ingress-mac"

    // Reply-ingress TLV is too short to contain the
    // Port ID type
    CfmPmPktAction_malformed_ingress_port_length_size CfmPmPktAction = "malformed-ingress-port-length-size"

    // Reply-ingress TLV has a zero Port ID length
    CfmPmPktAction_malformed_ingress_port_id_length CfmPmPktAction = "malformed-ingress-port-id-length"

    // Reply-ingress TLV is too short to contain the
    // Port ID
    CfmPmPktAction_malformed_ingress_port_id_size CfmPmPktAction = "malformed-ingress-port-id-size"

    // Reply-egress TLV is too short
    CfmPmPktAction_malformed_reply_egress_size CfmPmPktAction = "malformed-reply-egress-size"

    // Invalid egress-action value
    CfmPmPktAction_malformed_egress_action CfmPmPktAction = "malformed-egress-action"

    // Reply-egress TLV has invalid MAC address
    CfmPmPktAction_malformed_reply_egress_mac CfmPmPktAction = "malformed-reply-egress-mac"

    // Reply-egress TLV is too short to contain the
    // Port ID type
    CfmPmPktAction_malformed_egress_port_length_size CfmPmPktAction = "malformed-egress-port-length-size"

    // Reply-egress TLV has a zero Port ID length
    CfmPmPktAction_malformed_egress_port_id_length CfmPmPktAction = "malformed-egress-port-id-length"

    // Reply-egress TLV is too short to contain the
    // Port ID
    CfmPmPktAction_malformed_egress_port_id_size CfmPmPktAction = "malformed-egress-port-id-size"

    // LTM-Egress_ID TLV is too short
    CfmPmPktAction_malformed_ltm_egress_id_size CfmPmPktAction = "malformed-ltm-egress-id-size"

    // MEP-name TLV is too short
    CfmPmPktAction_malformed_mep_name_size CfmPmPktAction = "malformed-mep-name-size"

    // MEP-name TLV is too short to contain a MEP name
    CfmPmPktAction_malformed_mep_name_name_length CfmPmPktAction = "malformed-mep-name-name-length"

    // Additional-interface-status is too short
    CfmPmPktAction_malformed_additional_interface_status_size CfmPmPktAction = "malformed-additional-interface-status-size"

    // Invalid additional interface status
    CfmPmPktAction_malformed_additional_interface_status CfmPmPktAction = "malformed-additional-interface-status"

    // CCM has a zero CCM interval
    CfmPmPktAction_malformed_ccm_interval CfmPmPktAction = "malformed-ccm-interval"

    // CCM has a MAC-address MDID but the MDID is the
    // wrong length
    CfmPmPktAction_malformed_mdid_mac_address_length CfmPmPktAction = "malformed-mdid-mac-address-length"

    // CCM has an invalid MDID length
    CfmPmPktAction_malformed_mdid_length CfmPmPktAction = "malformed-mdid-length"

    // CCM has an invalid Short MA Name length
    CfmPmPktAction_malformed_sman_length CfmPmPktAction = "malformed-sman-length"

    // CCM has a VID or 16-bit Short MA Name but a
    // mismatched length
    CfmPmPktAction_malformed_sman2_byte_length CfmPmPktAction = "malformed-sman2-byte-length"

    // CCM has a VPNID Short MA Name but a mismatched
    // length
    CfmPmPktAction_malformed_sman_vpn_id_length CfmPmPktAction = "malformed-sman-vpn-id-length"

    // ELR has no ELR Reply TLVs
    CfmPmPktAction_malformed_elr_no_reply_tlv CfmPmPktAction = "malformed-elr-no-reply-tlv"

    // ELR Reply Egress TLVs not all adjacent
    CfmPmPktAction_malformed_separate_elr_reply_egress CfmPmPktAction = "malformed-separate-elr-reply-egress"

    // DCM has a multicast destination MAC
    CfmPmPktAction_malformed_dcm_destination_multicast CfmPmPktAction = "malformed-dcm-destination-multicast"

    // DCM is too short to contain an Embedded PDU
    CfmPmPktAction_malformed_dcm_embed_length CfmPmPktAction = "malformed-dcm-embed-length"

    // DCM Embedded PDU level does not match DCM level
    CfmPmPktAction_malformed_dcm_embed_level CfmPmPktAction = "malformed-dcm-embed-level"

    // DCM Embedded PDU version does not match DCM
    // version
    CfmPmPktAction_malformed_dcm_embed_version CfmPmPktAction = "malformed-dcm-embed-version"

    // Unknown ELR relay action
    CfmPmPktAction_malformed_elr_relay_action CfmPmPktAction = "malformed-elr-relay-action"

    // Reply Ingress TTL is not one greater than Reply
    // Egress TTL
    CfmPmPktAction_malformed_elr_tt_ls CfmPmPktAction = "malformed-elr-tt-ls"

    // Reply Ingress TTL present without ELR Reply
    // Ingress TLV
    CfmPmPktAction_malformed_elr_ttl_ingress CfmPmPktAction = "malformed-elr-ttl-ingress"

    // Reply Egress TTL present without ELR Reply
    // Egress TLV
    CfmPmPktAction_malformed_elr_ttl_egress CfmPmPktAction = "malformed-elr-ttl-egress"

    // ELM Destination MAC must not be unicast
    CfmPmPktAction_malformed_elm_destination_unicast CfmPmPktAction = "malformed-elm-destination-unicast"

    // ELM has no LTM Egress ID TLV
    CfmPmPktAction_malformed_elm_egress_id CfmPmPktAction = "malformed-elm-egress-id"

    // Embedded DCM OUI unrecognized
    CfmPmPktAction_malformed_dcm_embed_oui CfmPmPktAction = "malformed-dcm-embed-oui"

    // Embedded DCM Opcode is not ELM
    CfmPmPktAction_malformed_dcm_embed_opcode CfmPmPktAction = "malformed-dcm-embed-opcode"

    // ELM Constant Factor is zero
    CfmPmPktAction_malformed_elm_constant_zero CfmPmPktAction = "malformed-elm-constant-zero"

    // ELR Next-Hop Timeout is zero
    CfmPmPktAction_malformed_elr_timeout_zero CfmPmPktAction = "malformed-elr-timeout-zero"

    // Multiple Test TLVs found
    CfmPmPktAction_malformed_duplicate_test CfmPmPktAction = "malformed-duplicate-test"

    // Invalid source MAC address for DMM
    CfmPmPktAction_malformed_dmm_source_mac CfmPmPktAction = "malformed-dmm-source-mac"

    // Test TLV is too short
    CfmPmPktAction_malformed_test_size CfmPmPktAction = "malformed-test-size"

    // DMR has exactly one of its Rxf and Txb
    // timestamps unspecified
    CfmPmPktAction_malformed_dmr_time_stamps CfmPmPktAction = "malformed-dmr-time-stamps"

    // The format of one or more timestamps is invalid
    CfmPmPktAction_malformed_dm_time_stamp_fmt CfmPmPktAction = "malformed-dm-time-stamp-fmt"

    // AIS/LCK has invalid interval value (not 1
    // second or 1 minute)
    CfmPmPktAction_malformed_ais_interval CfmPmPktAction = "malformed-ais-interval"

    // Packet dropped due to interface being down
    CfmPmPktAction_filter_interface_down CfmPmPktAction = "filter-interface-down"

    // Packet dropped - not forwarded because
    // interface is in standby mode
    CfmPmPktAction_filter_forward_standby CfmPmPktAction = "filter-forward-standby"

    // CCM has an ICC-based format Short MA Name but a
    // mismatched length
    CfmPmPktAction_malformed_sman_icc_based_length CfmPmPktAction = "malformed-sman-icc-based-length"

    // Packet dropped - not forwarded in secondary HA
    // role
    CfmPmPktAction_filter_foward_issu_secondary CfmPmPktAction = "filter-foward-issu-secondary"

    // Packet dropped - not responded to because
    // interface is in standby mode
    CfmPmPktAction_filter_response_standby CfmPmPktAction = "filter-response-standby"

    // Packet dropped - not responded to in secondary
    // HA role
    CfmPmPktAction_filter_response_issu_secondary CfmPmPktAction = "filter-response-issu-secondary"
)

// SlaBucketSize represents Type of configuration of a bucket size
type SlaBucketSize string

const (
    // Bucket size is configured as buckets per probe
    SlaBucketSize_buckets_per_probe SlaBucketSize = "buckets-per-probe"

    // Bucket size is configured as probes per bucket
    SlaBucketSize_probes_per_bucket SlaBucketSize = "probes-per-bucket"
)

// CfmPmEltDelayModel represents operations
type CfmPmEltDelayModel string

const (
    // Not a valid delay model
    CfmPmEltDelayModel_delay_model_invalid CfmPmEltDelayModel = "delay-model-invalid"

    // Reply using logarithmic delay model
    CfmPmEltDelayModel_delay_model_logarithmic CfmPmEltDelayModel = "delay-model-logarithmic"

    // Reply using constant delay model
    CfmPmEltDelayModel_delay_model_constant CfmPmEltDelayModel = "delay-model-constant"
)

// CfmPmAisTransmit represents via a MIP or directly to a higher MEP
type CfmPmAisTransmit string

const (
    // AIS not transmitted
    CfmPmAisTransmit_transmit_none CfmPmAisTransmit = "transmit-none"

    // AIS transmitted via MIP
    CfmPmAisTransmit_transmit_ais CfmPmAisTransmit = "transmit-ais"

    // AIS signal passed directly to a higher MEP
    CfmPmAisTransmit_transmit_ais_direct CfmPmAisTransmit = "transmit-ais-direct"
)

// CfmPmElrRelayAction represents ELR relay action
type CfmPmElrRelayAction string

const (
    // Target Hit
    CfmPmElrRelayAction_elr_relay_hit CfmPmElrRelayAction = "elr-relay-hit"

    // Filtering database
    CfmPmElrRelayAction_elr_relay_fdb CfmPmElrRelayAction = "elr-relay-fdb"

    // Flood forwarded
    CfmPmElrRelayAction_elr_relay_flood CfmPmElrRelayAction = "elr-relay-flood"

    // Dropped
    CfmPmElrRelayAction_elr_relay_drop CfmPmElrRelayAction = "elr-relay-drop"
)

// CfmPmPortStatus represents Port status
type CfmPmPortStatus string

const (
    // Port is STP blocked
    CfmPmPortStatus_port_status_blocked CfmPmPortStatus = "port-status-blocked"

    // Port is up
    CfmPmPortStatus_port_status_up CfmPmPortStatus = "port-status-up"

    // Unknown port status
    CfmPmPortStatus_port_status_unknown CfmPmPortStatus = "port-status-unknown"
)

// CfmBagIwState represents CFM Interworking state
type CfmBagIwState string

const (
    // Interface is UP
    CfmBagIwState_interworking_up CfmBagIwState = "interworking-up"

    // Interface is in TEST mode
    CfmBagIwState_interworking_test CfmBagIwState = "interworking-test"
)

// CfmBagMdidFmt represents CFM MDID format
type CfmBagMdidFmt string

const (
    // MDID is explicity NULL
    CfmBagMdidFmt_mdid_null CfmBagMdidFmt = "mdid-null"

    // MDID is based on a DNS name
    CfmBagMdidFmt_mdid_dns_like CfmBagMdidFmt = "mdid-dns-like"

    // MDID is a (MAC address, integer) pair
    CfmBagMdidFmt_mdid_mac_address CfmBagMdidFmt = "mdid-mac-address"

    // MDID is a character string
    CfmBagMdidFmt_mdid_string CfmBagMdidFmt = "mdid-string"

    // Unknown MDID format
    CfmBagMdidFmt_mdid_unknown CfmBagMdidFmt = "mdid-unknown"
)

// CfmBagBdidFmt represents Bridge domain identifier format
type CfmBagBdidFmt string

const (
    // Invalid BDID identifier format
    CfmBagBdidFmt_invalid CfmBagBdidFmt = "invalid"

    // Identifier is a bridge domain ID
    CfmBagBdidFmt_bd_id CfmBagBdidFmt = "bd-id"

    // Identifier is a P2P cross-connect ID
    CfmBagBdidFmt_xc_p2p_id CfmBagBdidFmt = "xc-p2p-id"

    // Identifier is a MP2MP cross-connect ID
    CfmBagBdidFmt_xc_mp2mp_id CfmBagBdidFmt = "xc-mp2mp-id"

    // Identifier is a VLAN-aware flexible
    // cross-connect ID
    CfmBagBdidFmt_fxc_vlan_aware_id CfmBagBdidFmt = "fxc-vlan-aware-id"

    // Identifier is a VLAN-unaware flexible
    // cross-connect ID
    CfmBagBdidFmt_fxc_vlan_unaware_id CfmBagBdidFmt = "fxc-vlan-unaware-id"

    // Identifier is a maintenance association name
    CfmBagBdidFmt_down_only CfmBagBdidFmt = "down-only"
)

// CfmBagIssuRole represents CFM ISSU role
type CfmBagIssuRole string

const (
    // Unknown
    CfmBagIssuRole_unknown CfmBagIssuRole = "unknown"

    // Primary
    CfmBagIssuRole_primary CfmBagIssuRole = "primary"

    // Secondary
    CfmBagIssuRole_secondary CfmBagIssuRole = "secondary"
)

// CfmBagStpState represents CFM STP state
type CfmBagStpState string

const (
    // Interface is UP
    CfmBagStpState_stp_up CfmBagStpState = "stp-up"

    // Interface is STP-blocked
    CfmBagStpState_stp_blocked CfmBagStpState = "stp-blocked"

    // Unknown Interface STP state
    CfmBagStpState_stp_unknown CfmBagStpState = "stp-unknown"
)

// CfmBagMdLevel represents CFM level
type CfmBagMdLevel string

const (
    // CFM level 0
    CfmBagMdLevel_level0 CfmBagMdLevel = "level0"

    // CFM level 1
    CfmBagMdLevel_level1 CfmBagMdLevel = "level1"

    // CFM level 2
    CfmBagMdLevel_level2 CfmBagMdLevel = "level2"

    // CFM level 3
    CfmBagMdLevel_level3 CfmBagMdLevel = "level3"

    // CFM level 4
    CfmBagMdLevel_level4 CfmBagMdLevel = "level4"

    // CFM level 5
    CfmBagMdLevel_level5 CfmBagMdLevel = "level5"

    // CFM level 6
    CfmBagMdLevel_level6 CfmBagMdLevel = "level6"

    // CFM level 7
    CfmBagMdLevel_level7 CfmBagMdLevel = "level7"

    // Invalid CFM level
    CfmBagMdLevel_level_invalid CfmBagMdLevel = "level-invalid"
)

// SlaOperPacketPriority represents Priority scheme for packet priority
type SlaOperPacketPriority string

const (
    // Packet does not use any specified priority.
    SlaOperPacketPriority_priority_none SlaOperPacketPriority = "priority-none"

    // Packet uses a specified 3-bit COS priority
    // value.
    SlaOperPacketPriority_priority_cos SlaOperPacketPriority = "priority-cos"
)

// CfmBagAisInterval represents CFM AIS intervals
type CfmBagAisInterval string

const (
    // Invalid AIS interval
    CfmBagAisInterval_ais_interval_none CfmBagAisInterval = "ais-interval-none"

    // Interval of 1s
    CfmBagAisInterval_ais_interval1s CfmBagAisInterval = "ais-interval1s"

    // Interval of 1 min
    CfmBagAisInterval_ais_interval1m CfmBagAisInterval = "ais-interval1m"
)

// CfmPmRmepXcState represents Cross-check state of a peer MEP
type CfmPmRmepXcState string

const (
    // Cross-check OK
    CfmPmRmepXcState_cross_check_ok CfmPmRmepXcState = "cross-check-ok"

    // No CCMs received within loss time from peer MEP
    CfmPmRmepXcState_cross_check_missing CfmPmRmepXcState = "cross-check-missing"

    // CCMs received from peer MEP not marked for
    // cross-check
    CfmPmRmepXcState_cross_check_extra CfmPmRmepXcState = "cross-check-extra"
)

// CfmPmLtMode represents Type of Linktrace operation
type CfmPmLtMode string

const (
    // Basic IEEE 802.1ag Linktrace
    CfmPmLtMode_cfm_pm_lt_mode_basic CfmPmLtMode = "cfm-pm-lt-mode-basic"

    // Cisco Exploratory Linktrace
    CfmPmLtMode_cfm_pm_lt_mode_exploratory CfmPmLtMode = "cfm-pm-lt-mode-exploratory"
)

// CfmPmIntfStatus represents Interface status
type CfmPmIntfStatus string

const (
    // Interface is up
    CfmPmIntfStatus_interface_status_up CfmPmIntfStatus = "interface-status-up"

    // Interface is down
    CfmPmIntfStatus_interface_status_down CfmPmIntfStatus = "interface-status-down"

    // Interface is in testing mode
    CfmPmIntfStatus_interface_status_testing CfmPmIntfStatus = "interface-status-testing"

    // Unknown interface status
    CfmPmIntfStatus_interface_status_unknown CfmPmIntfStatus = "interface-status-unknown"

    // Interface is dormant
    CfmPmIntfStatus_interface_status_dormant CfmPmIntfStatus = "interface-status-dormant"

    // Interface status not found
    CfmPmIntfStatus_interface_status_not_present CfmPmIntfStatus = "interface-status-not-present"

    // Lower layer is down
    CfmPmIntfStatus_interface_status_lower_layer_down CfmPmIntfStatus = "interface-status-lower-layer-down"
)

// CfmBagDirection represents MEP direction
type CfmBagDirection string

const (
    // Up
    CfmBagDirection_direction_up CfmBagDirection = "direction-up"

    // Down
    CfmBagDirection_direction_down CfmBagDirection = "direction-down"

    // Invalid direction
    CfmBagDirection_direction_invalid CfmBagDirection = "direction-invalid"
)

// CfmPmEgressAction represents Egress action
type CfmPmEgressAction string

const (
    // OK
    CfmPmEgressAction_egress_ok CfmPmEgressAction = "egress-ok"

    // Down
    CfmPmEgressAction_egress_down CfmPmEgressAction = "egress-down"

    // STP Blocked
    CfmPmEgressAction_egress_blocked CfmPmEgressAction = "egress-blocked"

    // VID Blocked
    CfmPmEgressAction_egress_vid CfmPmEgressAction = "egress-vid"
)

// CfmPmElmReplyFilter represents operations
type CfmPmElmReplyFilter string

const (
    // Reply Filter not present
    CfmPmElmReplyFilter_reply_filter_not_present CfmPmElmReplyFilter = "reply-filter-not-present"

    // Reply from ports which are not MAC-pruned,
    // VID-pruned, or STP-blocked
    CfmPmElmReplyFilter_reply_filter_default CfmPmElmReplyFilter = "reply-filter-default"

    // Reply from ports which are not VID-pruned or
    // STP-blocked
    CfmPmElmReplyFilter_reply_filter_vlan_topology CfmPmElmReplyFilter = "reply-filter-vlan-topology"

    // Reply from ports which are not STP-blocked
    CfmPmElmReplyFilter_reply_filter_spanning_tree CfmPmElmReplyFilter = "reply-filter-spanning-tree"

    // Reply from all ports
    CfmPmElmReplyFilter_reply_filter_all_ports CfmPmElmReplyFilter = "reply-filter-all-ports"
)

// CfmAisDir represents Cfm ais dir
type CfmAisDir string

const (
    // Packets sent inward
    CfmAisDir_up CfmAisDir = "up"

    // Packets sent outward
    CfmAisDir_down CfmAisDir = "down"
)

// CfmPmAddlIntfStatus represents Additional interface status
type CfmPmAddlIntfStatus string

const (
    // Additional interface status unknown
    CfmPmAddlIntfStatus_unknown CfmPmAddlIntfStatus = "unknown"

    // Interface is explicitly shutdown in
    // configuration
    CfmPmAddlIntfStatus_administratively_down CfmPmAddlIntfStatus = "administratively-down"

    // Remote interface has exceeded its 802.3 Link
    // OAM error threshold
    CfmPmAddlIntfStatus_remote_excessive_errors CfmPmAddlIntfStatus = "remote-excessive-errors"

    // Local interface has exceeded its 802.3 Link OAM
    // error threshold
    CfmPmAddlIntfStatus_local_excessive_errors CfmPmAddlIntfStatus = "local-excessive-errors"
)

// CfmBagOpcode represents CFM Opcode
type CfmBagOpcode string

const (
    // Reserved
    CfmBagOpcode_reserved CfmBagOpcode = "reserved"

    // Continuity Check
    CfmBagOpcode_ccm CfmBagOpcode = "ccm"

    // Loopback Reply
    CfmBagOpcode_lbr CfmBagOpcode = "lbr"

    // Loopback Message
    CfmBagOpcode_lbm CfmBagOpcode = "lbm"

    // Linktrace Reply
    CfmBagOpcode_ltr CfmBagOpcode = "ltr"

    // Linktrace Message
    CfmBagOpcode_ltm CfmBagOpcode = "ltm"
)

// SlaOperTestPatternScheme represents Test pattern scheme for packet padding
type SlaOperTestPatternScheme string

const (
    // Packet is padded with a user-specified string
    SlaOperTestPatternScheme_hex SlaOperTestPatternScheme = "hex"

    // Packet is padded with a pseudo-random bit
    // sequence
    SlaOperTestPatternScheme_pseudo_random SlaOperTestPatternScheme = "pseudo-random"
)

// CfmPmMepFngState represents states
type CfmPmMepFngState string

const (
    // FNG in reset state
    CfmPmMepFngState_fng_reset CfmPmMepFngState = "fng-reset"

    // FNG has detected but not yet reported a defect
    CfmPmMepFngState_fng_defect CfmPmMepFngState = "fng-defect"

    // FNG is in the process of reporting a defect
    CfmPmMepFngState_fng_report_defect CfmPmMepFngState = "fng-report-defect"

    // FNG has reported a defect
    CfmPmMepFngState_fng_defect_reported CfmPmMepFngState = "fng-defect-reported"

    // No defect present, but the reset timer has not
    // yet expired
    CfmPmMepFngState_fng_defect_clearing CfmPmMepFngState = "fng-defect-clearing"
)

// CfmPmPortIdFmt represents Port ID format
type CfmPmPortIdFmt string

const (
    // Port ID is an interface alias
    CfmPmPortIdFmt_port_id_interface_alias CfmPmPortIdFmt = "port-id-interface-alias"

    // Port ID is a component name
    CfmPmPortIdFmt_port_id_port_component CfmPmPortIdFmt = "port-id-port-component"

    // Port ID is a MAC address
    CfmPmPortIdFmt_port_id_mac_address CfmPmPortIdFmt = "port-id-mac-address"

    // Port ID is a network address
    CfmPmPortIdFmt_port_id_network_address CfmPmPortIdFmt = "port-id-network-address"

    // Port ID is an interface name
    CfmPmPortIdFmt_port_id_interface_name CfmPmPortIdFmt = "port-id-interface-name"

    // Port ID is an agent name
    CfmPmPortIdFmt_port_id_agent_circuit_id CfmPmPortIdFmt = "port-id-agent-circuit-id"

    // Port ID is a local name
    CfmPmPortIdFmt_port_id_local CfmPmPortIdFmt = "port-id-local"

    // Port ID format unknown
    CfmPmPortIdFmt_port_id_unknown CfmPmPortIdFmt = "port-id-unknown"
)

// SlaOperBucket represents Type of SLA metric bucket
type SlaOperBucket string

const (
    // SLA metric bin
    SlaOperBucket_bucket_type_bins SlaOperBucket = "bucket-type-bins"

    // SLA metric sample
    SlaOperBucket_bucket_type_samples SlaOperBucket = "bucket-type-samples"
)

// SlaRecordableMetric represents Types of metrics that can be recorded by probes
type SlaRecordableMetric string

const (
    // Not a valid metric type
    SlaRecordableMetric_metric_invalid SlaRecordableMetric = "metric-invalid"

    // Round-trip Delay
    SlaRecordableMetric_metric_round_trip_delay SlaRecordableMetric = "metric-round-trip-delay"

    // One-way Delay (Source->Destination)
    SlaRecordableMetric_metric_one_way_delay_sd SlaRecordableMetric = "metric-one-way-delay-sd"

    // One-way Delay (Destination->Source)
    SlaRecordableMetric_metric_one_way_delay_ds SlaRecordableMetric = "metric-one-way-delay-ds"

    // Round-trip Jitter
    SlaRecordableMetric_metric_round_trip_jitter SlaRecordableMetric = "metric-round-trip-jitter"

    // One-way Jitter (Source->Destination)
    SlaRecordableMetric_metric_one_way_jitter_sd SlaRecordableMetric = "metric-one-way-jitter-sd"

    // One-way Jitter (Destination->Source)
    SlaRecordableMetric_metric_one_way_jitter_ds SlaRecordableMetric = "metric-one-way-jitter-ds"

    // One-way Frame Loss Ratio (Source->Destination)
    SlaRecordableMetric_metric_one_way_flr_sd SlaRecordableMetric = "metric-one-way-flr-sd"

    // One-way Frame Loss Ratio (Destination->Source)
    SlaRecordableMetric_metric_one_way_flr_ds SlaRecordableMetric = "metric-one-way-flr-ds"
)

// Cfm
// CFM operational data
type Cfm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node table for node-specific operational data.
    Nodes Cfm_Nodes

    // Global operational data.
    Global Cfm_Global
}

func (cfm *Cfm) GetFilter() yfilter.YFilter { return cfm.YFilter }

func (cfm *Cfm) SetFilter(yf yfilter.YFilter) { cfm.YFilter = yf }

func (cfm *Cfm) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "global" { return "Global" }
    return ""
}

func (cfm *Cfm) GetSegmentPath() string {
    return "Cisco-IOS-XR-ethernet-cfm-oper:cfm"
}

func (cfm *Cfm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &cfm.Nodes
    }
    if childYangName == "global" {
        return &cfm.Global
    }
    return nil
}

func (cfm *Cfm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &cfm.Nodes
    children["global"] = &cfm.Global
    return children
}

func (cfm *Cfm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cfm *Cfm) GetBundleName() string { return "cisco_ios_xr" }

func (cfm *Cfm) GetYangName() string { return "cfm" }

func (cfm *Cfm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cfm *Cfm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cfm *Cfm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cfm *Cfm) SetParent(parent types.Entity) { cfm.parent = parent }

func (cfm *Cfm) GetParent() types.Entity { return cfm.parent }

func (cfm *Cfm) GetParentYangName() string { return "Cisco-IOS-XR-ethernet-cfm-oper" }

// Cfm_Nodes
// Node table for node-specific operational data
type Cfm_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // Cfm_Nodes_Node.
    Node []Cfm_Nodes_Node
}

func (nodes *Cfm_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Cfm_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Cfm_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Cfm_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Cfm_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Cfm_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Cfm_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Cfm_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Cfm_Nodes) GetYangName() string { return "nodes" }

func (nodes *Cfm_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Cfm_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Cfm_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Cfm_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Cfm_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Cfm_Nodes) GetParentYangName() string { return "cfm" }

// Cfm_Nodes_Node
// Node-specific data for a particular node
type Cfm_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Interface AIS table.
    InterfaceAises Cfm_Nodes_Node_InterfaceAises

    // Interface Statistics table.
    InterfaceStatistics Cfm_Nodes_Node_InterfaceStatistics

    // Summary.
    Summary Cfm_Nodes_Node_Summary

    // CCMLearningDatabase table.
    CcmLearningDatabases Cfm_Nodes_Node_CcmLearningDatabases
}

func (node *Cfm_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Cfm_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Cfm_Nodes_Node) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "interface-aises" { return "InterfaceAises" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    if yname == "summary" { return "Summary" }
    if yname == "ccm-learning-databases" { return "CcmLearningDatabases" }
    return ""
}

func (node *Cfm_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
}

func (node *Cfm_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-aises" {
        return &node.InterfaceAises
    }
    if childYangName == "interface-statistics" {
        return &node.InterfaceStatistics
    }
    if childYangName == "summary" {
        return &node.Summary
    }
    if childYangName == "ccm-learning-databases" {
        return &node.CcmLearningDatabases
    }
    return nil
}

func (node *Cfm_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-aises"] = &node.InterfaceAises
    children["interface-statistics"] = &node.InterfaceStatistics
    children["summary"] = &node.Summary
    children["ccm-learning-databases"] = &node.CcmLearningDatabases
    return children
}

func (node *Cfm_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = node.Node
    return leafs
}

func (node *Cfm_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Cfm_Nodes_Node) GetYangName() string { return "node" }

func (node *Cfm_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Cfm_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Cfm_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Cfm_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Cfm_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Cfm_Nodes_Node) GetParentYangName() string { return "nodes" }

// Cfm_Nodes_Node_InterfaceAises
// Interface AIS table
type Cfm_Nodes_Node_InterfaceAises struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AIS statistics for a particular interface. The type is slice of
    // Cfm_Nodes_Node_InterfaceAises_InterfaceAis.
    InterfaceAis []Cfm_Nodes_Node_InterfaceAises_InterfaceAis
}

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetFilter() yfilter.YFilter { return interfaceAises.YFilter }

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) SetFilter(yf yfilter.YFilter) { interfaceAises.YFilter = yf }

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetGoName(yname string) string {
    if yname == "interface-ais" { return "InterfaceAis" }
    return ""
}

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetSegmentPath() string {
    return "interface-aises"
}

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-ais" {
        for _, c := range interfaceAises.InterfaceAis {
            if interfaceAises.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Nodes_Node_InterfaceAises_InterfaceAis{}
        interfaceAises.InterfaceAis = append(interfaceAises.InterfaceAis, child)
        return &interfaceAises.InterfaceAis[len(interfaceAises.InterfaceAis)-1]
    }
    return nil
}

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceAises.InterfaceAis {
        children[interfaceAises.InterfaceAis[i].GetSegmentPath()] = &interfaceAises.InterfaceAis[i]
    }
    return children
}

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetYangName() string { return "interface-aises" }

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) SetParent(parent types.Entity) { interfaceAises.parent = parent }

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetParent() types.Entity { return interfaceAises.parent }

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetParentYangName() string { return "node" }

// Cfm_Nodes_Node_InterfaceAises_InterfaceAis
// AIS statistics for a particular interface
type Cfm_Nodes_Node_InterfaceAises_InterfaceAis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. AIS Direction. The type is CfmAisDir.
    Direction interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // IM Interface state. The type is string.
    InterfaceState interface{}

    // Interface interworking state. The type is CfmBagIwState.
    InterworkingState interface{}

    // STP state. The type is CfmBagStpState.
    StpState interface{}

    // AIS statistics.
    Statistics Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics
}

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetFilter() yfilter.YFilter { return interfaceAis.YFilter }

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) SetFilter(yf yfilter.YFilter) { interfaceAis.YFilter = yf }

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "direction" { return "Direction" }
    if yname == "interface" { return "Interface" }
    if yname == "interface-state" { return "InterfaceState" }
    if yname == "interworking-state" { return "InterworkingState" }
    if yname == "stp-state" { return "StpState" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetSegmentPath() string {
    return "interface-ais" + "[interface-name='" + fmt.Sprintf("%v", interfaceAis.InterfaceName) + "']" + "[direction='" + fmt.Sprintf("%v", interfaceAis.Direction) + "']"
}

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &interfaceAis.Statistics
    }
    return nil
}

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &interfaceAis.Statistics
    return children
}

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceAis.InterfaceName
    leafs["direction"] = interfaceAis.Direction
    leafs["interface"] = interfaceAis.Interface
    leafs["interface-state"] = interfaceAis.InterfaceState
    leafs["interworking-state"] = interfaceAis.InterworkingState
    leafs["stp-state"] = interfaceAis.StpState
    return leafs
}

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetYangName() string { return "interface-ais" }

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) SetParent(parent types.Entity) { interfaceAis.parent = parent }

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetParent() types.Entity { return interfaceAis.parent }

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetParentYangName() string { return "interface-aises" }

// Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics
// AIS statistics
type Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Direction of AIS packets. The type is CfmBagDirection.
    Direction interface{}

    // Level of the lowest MEP transmitting AIS. The type is CfmBagMdLevel.
    LowestLevel interface{}

    // Level that AIS packets are transmitted on. The type is CfmBagMdLevel.
    TransmissionLevel interface{}

    // Interval at which AIS packets are transmitted. The type is
    // CfmBagAisInterval.
    TransmissionInterval interface{}

    // Total number of packets sent by the transmitting MEP. The type is
    // interface{} with range: 0..4294967295.
    SentPackets interface{}

    // Levels of other MEPs receiving AIS. The type is slice of CfmBagMdLevel.
    ViaLevel []interface{}

    // Defects detected.
    Defects Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects

    // Time elapsed since sending last started.
    LastStarted Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted
}

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "lowest-level" { return "LowestLevel" }
    if yname == "transmission-level" { return "TransmissionLevel" }
    if yname == "transmission-interval" { return "TransmissionInterval" }
    if yname == "sent-packets" { return "SentPackets" }
    if yname == "via-level" { return "ViaLevel" }
    if yname == "defects" { return "Defects" }
    if yname == "last-started" { return "LastStarted" }
    return ""
}

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "defects" {
        return &statistics.Defects
    }
    if childYangName == "last-started" {
        return &statistics.LastStarted
    }
    return nil
}

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["defects"] = &statistics.Defects
    children["last-started"] = &statistics.LastStarted
    return children
}

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = statistics.Direction
    leafs["lowest-level"] = statistics.LowestLevel
    leafs["transmission-level"] = statistics.TransmissionLevel
    leafs["transmission-interval"] = statistics.TransmissionInterval
    leafs["sent-packets"] = statistics.SentPackets
    leafs["via-level"] = statistics.ViaLevel
    return leafs
}

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetYangName() string { return "statistics" }

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetParentYangName() string { return "interface-ais" }

// Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects
// Defects detected
type Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AIS or LCK received. The type is bool.
    AisReceived interface{}

    // Number of peer MEPs that have timed out. The type is interface{} with
    // range: 0..4294967295.
    PeerMepsThatTimedOut interface{}

    // Number of missing peer MEPs. The type is interface{} with range:
    // 0..4294967295.
    Missing interface{}

    // Number of missing auto cross-check MEPs. The type is interface{} with
    // range: 0..4294967295.
    AutoMissing interface{}

    // Number of unexpected peer MEPs. The type is interface{} with range:
    // 0..4294967295.
    Unexpected interface{}

    // The local port or interface is down. The type is bool.
    LocalPortStatus interface{}

    // A peer port or interface is down. The type is bool.
    PeerPortStatus interface{}

    // Defects detected from remote MEPs.
    RemoteMepsDefects Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects
}

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetFilter() yfilter.YFilter { return defects.YFilter }

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) SetFilter(yf yfilter.YFilter) { defects.YFilter = yf }

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetGoName(yname string) string {
    if yname == "ais-received" { return "AisReceived" }
    if yname == "peer-meps-that-timed-out" { return "PeerMepsThatTimedOut" }
    if yname == "missing" { return "Missing" }
    if yname == "auto-missing" { return "AutoMissing" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "local-port-status" { return "LocalPortStatus" }
    if yname == "peer-port-status" { return "PeerPortStatus" }
    if yname == "remote-meps-defects" { return "RemoteMepsDefects" }
    return ""
}

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetSegmentPath() string {
    return "defects"
}

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-meps-defects" {
        return &defects.RemoteMepsDefects
    }
    return nil
}

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote-meps-defects"] = &defects.RemoteMepsDefects
    return children
}

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ais-received"] = defects.AisReceived
    leafs["peer-meps-that-timed-out"] = defects.PeerMepsThatTimedOut
    leafs["missing"] = defects.Missing
    leafs["auto-missing"] = defects.AutoMissing
    leafs["unexpected"] = defects.Unexpected
    leafs["local-port-status"] = defects.LocalPortStatus
    leafs["peer-port-status"] = defects.PeerPortStatus
    return leafs
}

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetBundleName() string { return "cisco_ios_xr" }

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetYangName() string { return "defects" }

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) SetParent(parent types.Entity) { defects.parent = parent }

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetParent() types.Entity { return defects.parent }

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetParentYangName() string { return "statistics" }

// Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects
// Defects detected from remote MEPs
type Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timed out (loss threshold exceeded). The type is bool.
    LossThresholdExceeded interface{}

    // Invalid level. The type is bool.
    InvalidLevel interface{}

    // Invalid MAID. The type is bool.
    InvalidMaid interface{}

    // Invalid CCM interval. The type is bool.
    InvalidCcmInterval interface{}

    // Loop detected (our MAC address received). The type is bool.
    ReceivedOurMac interface{}

    // Configuration Error (our MEP ID received). The type is bool.
    ReceivedOurMepId interface{}

    // Remote defection indication received. The type is bool.
    ReceivedRdi interface{}
}

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetFilter() yfilter.YFilter { return remoteMepsDefects.YFilter }

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) SetFilter(yf yfilter.YFilter) { remoteMepsDefects.YFilter = yf }

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetGoName(yname string) string {
    if yname == "loss-threshold-exceeded" { return "LossThresholdExceeded" }
    if yname == "invalid-level" { return "InvalidLevel" }
    if yname == "invalid-maid" { return "InvalidMaid" }
    if yname == "invalid-ccm-interval" { return "InvalidCcmInterval" }
    if yname == "received-our-mac" { return "ReceivedOurMac" }
    if yname == "received-our-mep-id" { return "ReceivedOurMepId" }
    if yname == "received-rdi" { return "ReceivedRdi" }
    return ""
}

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetSegmentPath() string {
    return "remote-meps-defects"
}

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["loss-threshold-exceeded"] = remoteMepsDefects.LossThresholdExceeded
    leafs["invalid-level"] = remoteMepsDefects.InvalidLevel
    leafs["invalid-maid"] = remoteMepsDefects.InvalidMaid
    leafs["invalid-ccm-interval"] = remoteMepsDefects.InvalidCcmInterval
    leafs["received-our-mac"] = remoteMepsDefects.ReceivedOurMac
    leafs["received-our-mep-id"] = remoteMepsDefects.ReceivedOurMepId
    leafs["received-rdi"] = remoteMepsDefects.ReceivedRdi
    return leafs
}

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetBundleName() string { return "cisco_ios_xr" }

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetYangName() string { return "remote-meps-defects" }

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) SetParent(parent types.Entity) { remoteMepsDefects.parent = parent }

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetParent() types.Entity { return remoteMepsDefects.parent }

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetParentYangName() string { return "defects" }

// Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted
// Time elapsed since sending last started
type Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetFilter() yfilter.YFilter { return lastStarted.YFilter }

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) SetFilter(yf yfilter.YFilter) { lastStarted.YFilter = yf }

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetSegmentPath() string {
    return "last-started"
}

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = lastStarted.Seconds
    leafs["nanoseconds"] = lastStarted.Nanoseconds
    return leafs
}

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetBundleName() string { return "cisco_ios_xr" }

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetYangName() string { return "last-started" }

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) SetParent(parent types.Entity) { lastStarted.parent = parent }

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetParent() types.Entity { return lastStarted.parent }

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetParentYangName() string { return "statistics" }

// Cfm_Nodes_Node_InterfaceStatistics
// Interface Statistics table
type Cfm_Nodes_Node_InterfaceStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Counters for a particular interface. The type is slice of
    // Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic.
    InterfaceStatistic []Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic
}

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetFilter() yfilter.YFilter { return interfaceStatistics.YFilter }

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) SetFilter(yf yfilter.YFilter) { interfaceStatistics.YFilter = yf }

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetGoName(yname string) string {
    if yname == "interface-statistic" { return "InterfaceStatistic" }
    return ""
}

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetSegmentPath() string {
    return "interface-statistics"
}

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-statistic" {
        for _, c := range interfaceStatistics.InterfaceStatistic {
            if interfaceStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic{}
        interfaceStatistics.InterfaceStatistic = append(interfaceStatistics.InterfaceStatistic, child)
        return &interfaceStatistics.InterfaceStatistic[len(interfaceStatistics.InterfaceStatistic)-1]
    }
    return nil
}

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceStatistics.InterfaceStatistic {
        children[interfaceStatistics.InterfaceStatistic[i].GetSegmentPath()] = &interfaceStatistics.InterfaceStatistic[i]
    }
    return children
}

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetYangName() string { return "interface-statistics" }

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) SetParent(parent types.Entity) { interfaceStatistics.parent = parent }

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetParent() types.Entity { return interfaceStatistics.parent }

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetParentYangName() string { return "node" }

// Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic
// Counters for a particular interface
type Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceXr interface{}

    // EFP statistics.
    Statistics Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics
}

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetFilter() yfilter.YFilter { return interfaceStatistic.YFilter }

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) SetFilter(yf yfilter.YFilter) { interfaceStatistic.YFilter = yf }

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "interface-xr" { return "InterfaceXr" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetSegmentPath() string {
    return "interface-statistic" + "[interface='" + fmt.Sprintf("%v", interfaceStatistic.Interface) + "']"
}

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &interfaceStatistic.Statistics
    }
    return nil
}

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &interfaceStatistic.Statistics
    return children
}

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = interfaceStatistic.Interface
    leafs["interface-xr"] = interfaceStatistic.InterfaceXr
    return leafs
}

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetYangName() string { return "interface-statistic" }

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) SetParent(parent types.Entity) { interfaceStatistic.parent = parent }

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetParent() types.Entity { return interfaceStatistic.parent }

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetParentYangName() string { return "interface-statistics" }

// Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics
// EFP statistics
type Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of malformed packets received at this EFP. The type is interface{}
    // with range: 0..18446744073709551615.
    MalformedPackets interface{}

    // Number of packets dropped at this EFP. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}

    // Opcode for last malformed packet. The type is CfmBagOpcode.
    LastMalformedOpcode interface{}

    // Reason last malformed packet was malformed. The type is CfmPmPktAction.
    LastMalformedReason interface{}
}

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetGoName(yname string) string {
    if yname == "malformed-packets" { return "MalformedPackets" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    if yname == "last-malformed-opcode" { return "LastMalformedOpcode" }
    if yname == "last-malformed-reason" { return "LastMalformedReason" }
    return ""
}

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["malformed-packets"] = statistics.MalformedPackets
    leafs["dropped-packets"] = statistics.DroppedPackets
    leafs["last-malformed-opcode"] = statistics.LastMalformedOpcode
    leafs["last-malformed-reason"] = statistics.LastMalformedReason
    return leafs
}

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetYangName() string { return "statistics" }

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetParentYangName() string { return "interface-statistic" }

// Cfm_Nodes_Node_Summary
// Summary
type Cfm_Nodes_Node_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of domains in the CFM database. The type is interface{} with
    // range: 0..4294967295.
    Domains interface{}

    // The number of services in the CFM database. The type is interface{} with
    // range: 0..4294967295.
    Services interface{}

    // The combined rate of CCMs on this card. The type is interface{} with range:
    // 0..4294967295.
    CcmRate interface{}

    // The number of local MEPs in the CFM database. The type is interface{} with
    // range: 0..4294967295.
    LocalMeps interface{}

    // The number of operational local MEPs. The type is interface{} with range:
    // 0..4294967295.
    OperationalLocalMeps interface{}

    // The number of down-MEPs. The type is interface{} with range: 0..4294967295.
    DownMeps interface{}

    // The number of up-MEPs. The type is interface{} with range: 0..4294967295.
    UpMeps interface{}

    // The number of MEPs for which CCM processing has been offloaded. The type is
    // interface{} with range: 0..4294967295.
    Offloaded interface{}

    // The number of MEPs offloaded with CCMs at 3.3ms intervals. The type is
    // interface{} with range: 0..4294967295.
    OffloadedAt33Ms interface{}

    // The number of MEPs offloaded with CCMs at 10ms intervals. The type is
    // interface{} with range: 0..4294967295.
    OffloadedAt10Ms interface{}

    // The number of local MEPs disabled due to configuration errors. The type is
    // interface{} with range: 0..4294967295.
    DisabledMisconfigured interface{}

    // The number of local MEPs disabled due to lack of resources. The type is
    // interface{} with range: 0..4294967295.
    DisabledOutOfResources interface{}

    // The number of local MEPs disabled due to operational errors. The type is
    // interface{} with range: 0..4294967295.
    DisabledOperationalError interface{}

    // The number of peer MEPs. The type is interface{} with range: 0..4294967295.
    PeerMeps interface{}

    // The number of operational peer MEPs recorded in the CFM database. The type
    // is interface{} with range: 0..4294967295.
    OperationalPeerMeps interface{}

    // The number of peer MEPs with defects. The type is interface{} with range:
    // 0..4294967295.
    PeerMepsWithDefects interface{}

    // The number of peer MEPs without defects. The type is interface{} with
    // range: 0..4294967295.
    PeerMepsWithoutDefects interface{}

    // The number of peer MEPs that have timed out. The type is interface{} with
    // range: 0..4294967295.
    PeerMepsTimedOut interface{}

    // The number of MIPs. The type is interface{} with range: 0..4294967295.
    Mips interface{}

    // The number of interfaces running CFM. The type is interface{} with range:
    // 0..4294967295.
    Interfaces interface{}

    // Number or bridge domains and crossconnects. The type is interface{} with
    // range: 0..4294967295.
    BridgeDomainsAndXconnects interface{}

    // Number of traceroute cache entries. The type is interface{} with range:
    // 0..4294967295.
    TracerouteCacheEntries interface{}

    // Number of traceroute cache replies. The type is interface{} with range:
    // 0..4294967295.
    TracerouteCacheReplies interface{}

    // Number of entries in the CCM learning database. The type is interface{}
    // with range: 0..4294967295.
    CcmLearningDbEntries interface{}

    // ISSU Role of CFM-D, if any. The type is CfmBagIssuRole.
    IssuRole interface{}

    // Number of BNM Enabled Links. The type is interface{} with range:
    // 0..4294967295.
    BnmEnabledLinks interface{}
}

func (summary *Cfm_Nodes_Node_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Cfm_Nodes_Node_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Cfm_Nodes_Node_Summary) GetGoName(yname string) string {
    if yname == "domains" { return "Domains" }
    if yname == "services" { return "Services" }
    if yname == "ccm-rate" { return "CcmRate" }
    if yname == "local-meps" { return "LocalMeps" }
    if yname == "operational-local-meps" { return "OperationalLocalMeps" }
    if yname == "down-meps" { return "DownMeps" }
    if yname == "up-meps" { return "UpMeps" }
    if yname == "offloaded" { return "Offloaded" }
    if yname == "offloaded-at3-3ms" { return "OffloadedAt33Ms" }
    if yname == "offloaded-at10ms" { return "OffloadedAt10Ms" }
    if yname == "disabled-misconfigured" { return "DisabledMisconfigured" }
    if yname == "disabled-out-of-resources" { return "DisabledOutOfResources" }
    if yname == "disabled-operational-error" { return "DisabledOperationalError" }
    if yname == "peer-meps" { return "PeerMeps" }
    if yname == "operational-peer-meps" { return "OperationalPeerMeps" }
    if yname == "peer-meps-with-defects" { return "PeerMepsWithDefects" }
    if yname == "peer-meps-without-defects" { return "PeerMepsWithoutDefects" }
    if yname == "peer-meps-timed-out" { return "PeerMepsTimedOut" }
    if yname == "mips" { return "Mips" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "bridge-domains-and-xconnects" { return "BridgeDomainsAndXconnects" }
    if yname == "traceroute-cache-entries" { return "TracerouteCacheEntries" }
    if yname == "traceroute-cache-replies" { return "TracerouteCacheReplies" }
    if yname == "ccm-learning-db-entries" { return "CcmLearningDbEntries" }
    if yname == "issu-role" { return "IssuRole" }
    if yname == "bnm-enabled-links" { return "BnmEnabledLinks" }
    return ""
}

func (summary *Cfm_Nodes_Node_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Cfm_Nodes_Node_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Cfm_Nodes_Node_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Cfm_Nodes_Node_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domains"] = summary.Domains
    leafs["services"] = summary.Services
    leafs["ccm-rate"] = summary.CcmRate
    leafs["local-meps"] = summary.LocalMeps
    leafs["operational-local-meps"] = summary.OperationalLocalMeps
    leafs["down-meps"] = summary.DownMeps
    leafs["up-meps"] = summary.UpMeps
    leafs["offloaded"] = summary.Offloaded
    leafs["offloaded-at3-3ms"] = summary.OffloadedAt33Ms
    leafs["offloaded-at10ms"] = summary.OffloadedAt10Ms
    leafs["disabled-misconfigured"] = summary.DisabledMisconfigured
    leafs["disabled-out-of-resources"] = summary.DisabledOutOfResources
    leafs["disabled-operational-error"] = summary.DisabledOperationalError
    leafs["peer-meps"] = summary.PeerMeps
    leafs["operational-peer-meps"] = summary.OperationalPeerMeps
    leafs["peer-meps-with-defects"] = summary.PeerMepsWithDefects
    leafs["peer-meps-without-defects"] = summary.PeerMepsWithoutDefects
    leafs["peer-meps-timed-out"] = summary.PeerMepsTimedOut
    leafs["mips"] = summary.Mips
    leafs["interfaces"] = summary.Interfaces
    leafs["bridge-domains-and-xconnects"] = summary.BridgeDomainsAndXconnects
    leafs["traceroute-cache-entries"] = summary.TracerouteCacheEntries
    leafs["traceroute-cache-replies"] = summary.TracerouteCacheReplies
    leafs["ccm-learning-db-entries"] = summary.CcmLearningDbEntries
    leafs["issu-role"] = summary.IssuRole
    leafs["bnm-enabled-links"] = summary.BnmEnabledLinks
    return leafs
}

func (summary *Cfm_Nodes_Node_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Cfm_Nodes_Node_Summary) GetYangName() string { return "summary" }

func (summary *Cfm_Nodes_Node_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Cfm_Nodes_Node_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Cfm_Nodes_Node_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Cfm_Nodes_Node_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Cfm_Nodes_Node_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Cfm_Nodes_Node_Summary) GetParentYangName() string { return "node" }

// Cfm_Nodes_Node_CcmLearningDatabases
// CCMLearningDatabase table
type Cfm_Nodes_Node_CcmLearningDatabases struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CCM Learning Database entry. The type is slice of
    // Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase.
    CcmLearningDatabase []Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase
}

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetFilter() yfilter.YFilter { return ccmLearningDatabases.YFilter }

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) SetFilter(yf yfilter.YFilter) { ccmLearningDatabases.YFilter = yf }

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetGoName(yname string) string {
    if yname == "ccm-learning-database" { return "CcmLearningDatabase" }
    return ""
}

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetSegmentPath() string {
    return "ccm-learning-databases"
}

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ccm-learning-database" {
        for _, c := range ccmLearningDatabases.CcmLearningDatabase {
            if ccmLearningDatabases.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase{}
        ccmLearningDatabases.CcmLearningDatabase = append(ccmLearningDatabases.CcmLearningDatabase, child)
        return &ccmLearningDatabases.CcmLearningDatabase[len(ccmLearningDatabases.CcmLearningDatabase)-1]
    }
    return nil
}

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ccmLearningDatabases.CcmLearningDatabase {
        children[ccmLearningDatabases.CcmLearningDatabase[i].GetSegmentPath()] = &ccmLearningDatabases.CcmLearningDatabase[i]
    }
    return children
}

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetBundleName() string { return "cisco_ios_xr" }

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetYangName() string { return "ccm-learning-databases" }

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) SetParent(parent types.Entity) { ccmLearningDatabases.parent = parent }

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetParent() types.Entity { return ccmLearningDatabases.parent }

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetParentYangName() string { return "node" }

// Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase
// CCM Learning Database entry
type Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Maintenance domain name. The type is string.
    DomainXr interface{}

    // Maintenance level. The type is CfmBagMdLevel.
    Level interface{}

    // Maintenance association name. The type is string.
    ServiceXr interface{}

    // Source MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SourceMacAddress interface{}

    // The XID of the ingress interface for the CCM. The type is interface{} with
    // range: 0..4294967295.
    IngressInterface interface{}

    // The XID is stale and may have been reused for a different interface. The
    // type is bool.
    Stale interface{}

    // String representation of the Bridge Domain or Cross-Connect associated with
    // the ingress XID. The type is string.
    IngressInterfaceString interface{}
}

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetFilter() yfilter.YFilter { return ccmLearningDatabase.YFilter }

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) SetFilter(yf yfilter.YFilter) { ccmLearningDatabase.YFilter = yf }

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "service" { return "Service" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "domain-xr" { return "DomainXr" }
    if yname == "level" { return "Level" }
    if yname == "service-xr" { return "ServiceXr" }
    if yname == "source-mac-address" { return "SourceMacAddress" }
    if yname == "ingress-interface" { return "IngressInterface" }
    if yname == "stale" { return "Stale" }
    if yname == "ingress-interface-string" { return "IngressInterfaceString" }
    return ""
}

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetSegmentPath() string {
    return "ccm-learning-database" + "[domain='" + fmt.Sprintf("%v", ccmLearningDatabase.Domain) + "']" + "[service='" + fmt.Sprintf("%v", ccmLearningDatabase.Service) + "']" + "[mac-address='" + fmt.Sprintf("%v", ccmLearningDatabase.MacAddress) + "']"
}

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = ccmLearningDatabase.Domain
    leafs["service"] = ccmLearningDatabase.Service
    leafs["mac-address"] = ccmLearningDatabase.MacAddress
    leafs["domain-xr"] = ccmLearningDatabase.DomainXr
    leafs["level"] = ccmLearningDatabase.Level
    leafs["service-xr"] = ccmLearningDatabase.ServiceXr
    leafs["source-mac-address"] = ccmLearningDatabase.SourceMacAddress
    leafs["ingress-interface"] = ccmLearningDatabase.IngressInterface
    leafs["stale"] = ccmLearningDatabase.Stale
    leafs["ingress-interface-string"] = ccmLearningDatabase.IngressInterfaceString
    return leafs
}

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetBundleName() string { return "cisco_ios_xr" }

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetYangName() string { return "ccm-learning-database" }

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) SetParent(parent types.Entity) { ccmLearningDatabase.parent = parent }

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetParent() types.Entity { return ccmLearningDatabase.parent }

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetParentYangName() string { return "ccm-learning-databases" }

// Cfm_Global
// Global operational data
type Cfm_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Incomplete Traceroute table.
    IncompleteTraceroutes Cfm_Global_IncompleteTraceroutes

    // Maintenance Points table.
    MaintenancePoints Cfm_Global_MaintenancePoints

    // Global configuration errors table.
    GlobalConfigurationErrors Cfm_Global_GlobalConfigurationErrors

    // MEP configuration errors table.
    MepConfigurationErrors Cfm_Global_MepConfigurationErrors

    // Traceroute Cache table.
    TracerouteCaches Cfm_Global_TracerouteCaches

    // Local MEPs table.
    LocalMeps Cfm_Global_LocalMeps

    // Peer MEPs table Version 2.
    PeerMePv2S Cfm_Global_PeerMePv2S
}

func (global *Cfm_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *Cfm_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *Cfm_Global) GetGoName(yname string) string {
    if yname == "incomplete-traceroutes" { return "IncompleteTraceroutes" }
    if yname == "maintenance-points" { return "MaintenancePoints" }
    if yname == "global-configuration-errors" { return "GlobalConfigurationErrors" }
    if yname == "mep-configuration-errors" { return "MepConfigurationErrors" }
    if yname == "traceroute-caches" { return "TracerouteCaches" }
    if yname == "local-meps" { return "LocalMeps" }
    if yname == "peer-me-pv2s" { return "PeerMePv2S" }
    return ""
}

func (global *Cfm_Global) GetSegmentPath() string {
    return "global"
}

func (global *Cfm_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "incomplete-traceroutes" {
        return &global.IncompleteTraceroutes
    }
    if childYangName == "maintenance-points" {
        return &global.MaintenancePoints
    }
    if childYangName == "global-configuration-errors" {
        return &global.GlobalConfigurationErrors
    }
    if childYangName == "mep-configuration-errors" {
        return &global.MepConfigurationErrors
    }
    if childYangName == "traceroute-caches" {
        return &global.TracerouteCaches
    }
    if childYangName == "local-meps" {
        return &global.LocalMeps
    }
    if childYangName == "peer-me-pv2s" {
        return &global.PeerMePv2S
    }
    return nil
}

func (global *Cfm_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["incomplete-traceroutes"] = &global.IncompleteTraceroutes
    children["maintenance-points"] = &global.MaintenancePoints
    children["global-configuration-errors"] = &global.GlobalConfigurationErrors
    children["mep-configuration-errors"] = &global.MepConfigurationErrors
    children["traceroute-caches"] = &global.TracerouteCaches
    children["local-meps"] = &global.LocalMeps
    children["peer-me-pv2s"] = &global.PeerMePv2S
    return children
}

func (global *Cfm_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (global *Cfm_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *Cfm_Global) GetYangName() string { return "global" }

func (global *Cfm_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *Cfm_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *Cfm_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *Cfm_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *Cfm_Global) GetParent() types.Entity { return global.parent }

func (global *Cfm_Global) GetParentYangName() string { return "cfm" }

// Cfm_Global_IncompleteTraceroutes
// Incomplete Traceroute table
type Cfm_Global_IncompleteTraceroutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a traceroute operation that has not yet timed out. The
    // type is slice of Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute.
    IncompleteTraceroute []Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute
}

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetFilter() yfilter.YFilter { return incompleteTraceroutes.YFilter }

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) SetFilter(yf yfilter.YFilter) { incompleteTraceroutes.YFilter = yf }

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetGoName(yname string) string {
    if yname == "incomplete-traceroute" { return "IncompleteTraceroute" }
    return ""
}

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetSegmentPath() string {
    return "incomplete-traceroutes"
}

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "incomplete-traceroute" {
        for _, c := range incompleteTraceroutes.IncompleteTraceroute {
            if incompleteTraceroutes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute{}
        incompleteTraceroutes.IncompleteTraceroute = append(incompleteTraceroutes.IncompleteTraceroute, child)
        return &incompleteTraceroutes.IncompleteTraceroute[len(incompleteTraceroutes.IncompleteTraceroute)-1]
    }
    return nil
}

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range incompleteTraceroutes.IncompleteTraceroute {
        children[incompleteTraceroutes.IncompleteTraceroute[i].GetSegmentPath()] = &incompleteTraceroutes.IncompleteTraceroute[i]
    }
    return children
}

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetBundleName() string { return "cisco_ios_xr" }

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetYangName() string { return "incomplete-traceroutes" }

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) SetParent(parent types.Entity) { incompleteTraceroutes.parent = parent }

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetParent() types.Entity { return incompleteTraceroutes.parent }

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetParentYangName() string { return "global" }

// Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute
// Information about a traceroute operation that
// has not yet timed out
type Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. MEP ID. The type is interface{} with range:
    // 1..8191.
    MepId interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}

    // This attribute is a key. Transaction ID. The type is interface{} with
    // range: -2147483648..2147483647.
    TransactionId interface{}

    // Time (in seconds) before the traceroute completes. The type is interface{}
    // with range: 0..18446744073709551615. Units are second.
    TimeLeft interface{}

    // Information about the traceroute operation.
    TracerouteInformation Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation
}

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetFilter() yfilter.YFilter { return incompleteTraceroute.YFilter }

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) SetFilter(yf yfilter.YFilter) { incompleteTraceroute.YFilter = yf }

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "service" { return "Service" }
    if yname == "mep-id" { return "MepId" }
    if yname == "interface" { return "Interface" }
    if yname == "transaction-id" { return "TransactionId" }
    if yname == "time-left" { return "TimeLeft" }
    if yname == "traceroute-information" { return "TracerouteInformation" }
    return ""
}

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetSegmentPath() string {
    return "incomplete-traceroute" + "[domain='" + fmt.Sprintf("%v", incompleteTraceroute.Domain) + "']" + "[service='" + fmt.Sprintf("%v", incompleteTraceroute.Service) + "']" + "[mep-id='" + fmt.Sprintf("%v", incompleteTraceroute.MepId) + "']" + "[interface='" + fmt.Sprintf("%v", incompleteTraceroute.Interface) + "']" + "[transaction-id='" + fmt.Sprintf("%v", incompleteTraceroute.TransactionId) + "']"
}

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traceroute-information" {
        return &incompleteTraceroute.TracerouteInformation
    }
    return nil
}

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traceroute-information"] = &incompleteTraceroute.TracerouteInformation
    return children
}

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = incompleteTraceroute.Domain
    leafs["service"] = incompleteTraceroute.Service
    leafs["mep-id"] = incompleteTraceroute.MepId
    leafs["interface"] = incompleteTraceroute.Interface
    leafs["transaction-id"] = incompleteTraceroute.TransactionId
    leafs["time-left"] = incompleteTraceroute.TimeLeft
    return leafs
}

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetBundleName() string { return "cisco_ios_xr" }

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetYangName() string { return "incomplete-traceroute" }

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) SetParent(parent types.Entity) { incompleteTraceroute.parent = parent }

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetParent() types.Entity { return incompleteTraceroute.parent }

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetParentYangName() string { return "incomplete-traceroutes" }

// Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation
// Information about the traceroute operation
type Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maintenance domain name. The type is string.
    Domain interface{}

    // Service name. The type is string.
    Service interface{}

    // Maintenance level. The type is CfmBagMdLevel.
    Level interface{}

    // Source MEP ID. The type is interface{} with range: 0..65535.
    SourceMepId interface{}

    // Source interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Source MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SourceMacAddress interface{}

    // Target MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    TargetMacAddress interface{}

    // Directed MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DirectedMacAddress interface{}

    // Target MEP ID. The type is interface{} with range: 0..65535.
    TargetMepId interface{}

    // Timestamp of initiation time (seconds). The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    Timestamp interface{}

    // Time to live. The type is interface{} with range: 0..255.
    Ttl interface{}

    // Transaction ID. The type is interface{} with range: 0..4294967295.
    TransactionId interface{}

    // Options affecting traceroute behavior.
    Options Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options
}

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetFilter() yfilter.YFilter { return tracerouteInformation.YFilter }

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) SetFilter(yf yfilter.YFilter) { tracerouteInformation.YFilter = yf }

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "service" { return "Service" }
    if yname == "level" { return "Level" }
    if yname == "source-mep-id" { return "SourceMepId" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "source-mac-address" { return "SourceMacAddress" }
    if yname == "target-mac-address" { return "TargetMacAddress" }
    if yname == "directed-mac-address" { return "DirectedMacAddress" }
    if yname == "target-mep-id" { return "TargetMepId" }
    if yname == "timestamp" { return "Timestamp" }
    if yname == "ttl" { return "Ttl" }
    if yname == "transaction-id" { return "TransactionId" }
    if yname == "options" { return "Options" }
    return ""
}

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetSegmentPath() string {
    return "traceroute-information"
}

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "options" {
        return &tracerouteInformation.Options
    }
    return nil
}

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["options"] = &tracerouteInformation.Options
    return children
}

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = tracerouteInformation.Domain
    leafs["service"] = tracerouteInformation.Service
    leafs["level"] = tracerouteInformation.Level
    leafs["source-mep-id"] = tracerouteInformation.SourceMepId
    leafs["source-interface"] = tracerouteInformation.SourceInterface
    leafs["source-mac-address"] = tracerouteInformation.SourceMacAddress
    leafs["target-mac-address"] = tracerouteInformation.TargetMacAddress
    leafs["directed-mac-address"] = tracerouteInformation.DirectedMacAddress
    leafs["target-mep-id"] = tracerouteInformation.TargetMepId
    leafs["timestamp"] = tracerouteInformation.Timestamp
    leafs["ttl"] = tracerouteInformation.Ttl
    leafs["transaction-id"] = tracerouteInformation.TransactionId
    return leafs
}

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetBundleName() string { return "cisco_ios_xr" }

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetYangName() string { return "traceroute-information" }

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) SetParent(parent types.Entity) { tracerouteInformation.parent = parent }

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetParent() types.Entity { return tracerouteInformation.parent }

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetParentYangName() string { return "incomplete-traceroute" }

// Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options
// Options affecting traceroute behavior
type Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mode. The type is CfmPmLtMode.
    Mode interface{}

    // Options for a basic IEEE 802.1ag Linktrace.
    BasicOptions Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions

    // Options for an Exploratory Linktrace.
    ExploratoryOptions Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions
}

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetFilter() yfilter.YFilter { return options.YFilter }

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) SetFilter(yf yfilter.YFilter) { options.YFilter = yf }

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    if yname == "basic-options" { return "BasicOptions" }
    if yname == "exploratory-options" { return "ExploratoryOptions" }
    return ""
}

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetSegmentPath() string {
    return "options"
}

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-options" {
        return &options.BasicOptions
    }
    if childYangName == "exploratory-options" {
        return &options.ExploratoryOptions
    }
    return nil
}

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-options"] = &options.BasicOptions
    children["exploratory-options"] = &options.ExploratoryOptions
    return children
}

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode"] = options.Mode
    return leafs
}

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetBundleName() string { return "cisco_ios_xr" }

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetYangName() string { return "options" }

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) SetParent(parent types.Entity) { options.parent = parent }

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetParent() types.Entity { return options.parent }

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetParentYangName() string { return "traceroute-information" }

// Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions
// Options for a basic IEEE 802.1ag Linktrace
type Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Traceroute was initiated automatically. The type is bool.
    IsAuto interface{}

    // Only use the Filtering Database for forwarding lookups. The type is bool.
    FdbOnly interface{}
}

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetFilter() yfilter.YFilter { return basicOptions.YFilter }

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) SetFilter(yf yfilter.YFilter) { basicOptions.YFilter = yf }

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetGoName(yname string) string {
    if yname == "is-auto" { return "IsAuto" }
    if yname == "fdb-only" { return "FdbOnly" }
    return ""
}

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetSegmentPath() string {
    return "basic-options"
}

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-auto"] = basicOptions.IsAuto
    leafs["fdb-only"] = basicOptions.FdbOnly
    return leafs
}

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetBundleName() string { return "cisco_ios_xr" }

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetYangName() string { return "basic-options" }

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) SetParent(parent types.Entity) { basicOptions.parent = parent }

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetParent() types.Entity { return basicOptions.parent }

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetParentYangName() string { return "options" }

// Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions
// Options for an Exploratory Linktrace
type Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Delay model for delay calculations. The type is CfmPmEltDelayModel.
    DelayModel interface{}

    // Constant Factor for delay calculations. The type is interface{} with range:
    // 0..4294967295.
    DelayConstantFactor interface{}

    // Reply Filtering mode used by responders. The type is CfmPmElmReplyFilter.
    ReplyFilter interface{}
}

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetFilter() yfilter.YFilter { return exploratoryOptions.YFilter }

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) SetFilter(yf yfilter.YFilter) { exploratoryOptions.YFilter = yf }

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetGoName(yname string) string {
    if yname == "delay-model" { return "DelayModel" }
    if yname == "delay-constant-factor" { return "DelayConstantFactor" }
    if yname == "reply-filter" { return "ReplyFilter" }
    return ""
}

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetSegmentPath() string {
    return "exploratory-options"
}

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["delay-model"] = exploratoryOptions.DelayModel
    leafs["delay-constant-factor"] = exploratoryOptions.DelayConstantFactor
    leafs["reply-filter"] = exploratoryOptions.ReplyFilter
    return leafs
}

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetBundleName() string { return "cisco_ios_xr" }

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetYangName() string { return "exploratory-options" }

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) SetParent(parent types.Entity) { exploratoryOptions.parent = parent }

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetParent() types.Entity { return exploratoryOptions.parent }

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetParentYangName() string { return "options" }

// Cfm_Global_MaintenancePoints
// Maintenance Points table
type Cfm_Global_MaintenancePoints struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular Maintenance Point. The type is slice of
    // Cfm_Global_MaintenancePoints_MaintenancePoint.
    MaintenancePoint []Cfm_Global_MaintenancePoints_MaintenancePoint
}

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetFilter() yfilter.YFilter { return maintenancePoints.YFilter }

func (maintenancePoints *Cfm_Global_MaintenancePoints) SetFilter(yf yfilter.YFilter) { maintenancePoints.YFilter = yf }

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetGoName(yname string) string {
    if yname == "maintenance-point" { return "MaintenancePoint" }
    return ""
}

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetSegmentPath() string {
    return "maintenance-points"
}

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "maintenance-point" {
        for _, c := range maintenancePoints.MaintenancePoint {
            if maintenancePoints.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_MaintenancePoints_MaintenancePoint{}
        maintenancePoints.MaintenancePoint = append(maintenancePoints.MaintenancePoint, child)
        return &maintenancePoints.MaintenancePoint[len(maintenancePoints.MaintenancePoint)-1]
    }
    return nil
}

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range maintenancePoints.MaintenancePoint {
        children[maintenancePoints.MaintenancePoint[i].GetSegmentPath()] = &maintenancePoints.MaintenancePoint[i]
    }
    return children
}

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetBundleName() string { return "cisco_ios_xr" }

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetYangName() string { return "maintenance-points" }

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maintenancePoints *Cfm_Global_MaintenancePoints) SetParent(parent types.Entity) { maintenancePoints.parent = parent }

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetParent() types.Entity { return maintenancePoints.parent }

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetParentYangName() string { return "global" }

// Cfm_Global_MaintenancePoints_MaintenancePoint
// Information about a particular Maintenance
// Point
type Cfm_Global_MaintenancePoints_MaintenancePoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}

    // MEP error flag. The type is bool.
    MepHasError interface{}

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Maintenance Point.
    MaintenancePoint Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetFilter() yfilter.YFilter { return maintenancePoint.YFilter }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) SetFilter(yf yfilter.YFilter) { maintenancePoint.YFilter = yf }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "service" { return "Service" }
    if yname == "interface" { return "Interface" }
    if yname == "mep-has-error" { return "MepHasError" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "maintenance-point" { return "MaintenancePoint" }
    return ""
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetSegmentPath() string {
    return "maintenance-point" + "[domain='" + fmt.Sprintf("%v", maintenancePoint.Domain) + "']" + "[service='" + fmt.Sprintf("%v", maintenancePoint.Service) + "']" + "[interface='" + fmt.Sprintf("%v", maintenancePoint.Interface) + "']"
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "maintenance-point" {
        return &maintenancePoint.MaintenancePoint
    }
    return nil
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["maintenance-point"] = &maintenancePoint.MaintenancePoint
    return children
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = maintenancePoint.Domain
    leafs["service"] = maintenancePoint.Service
    leafs["interface"] = maintenancePoint.Interface
    leafs["mep-has-error"] = maintenancePoint.MepHasError
    leafs["mac-address"] = maintenancePoint.MacAddress
    return leafs
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetBundleName() string { return "cisco_ios_xr" }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetYangName() string { return "maintenance-point" }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) SetParent(parent types.Entity) { maintenancePoint.parent = parent }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetParent() types.Entity { return maintenancePoint.parent }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetParentYangName() string { return "maintenance-points" }

// Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint
// Maintenance Point
type Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain name. The type is string.
    DomainName interface{}

    // Domain level. The type is CfmBagMdLevel.
    Level interface{}

    // Service name. The type is string.
    ServiceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Type of Maintenance Point. The type is CfmMaMpVariety.
    MaintenancePointType interface{}

    // MEP ID. The type is interface{} with range: 0..65535.
    MepId interface{}
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetFilter() yfilter.YFilter { return maintenancePoint.YFilter }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) SetFilter(yf yfilter.YFilter) { maintenancePoint.YFilter = yf }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetGoName(yname string) string {
    if yname == "domain-name" { return "DomainName" }
    if yname == "level" { return "Level" }
    if yname == "service-name" { return "ServiceName" }
    if yname == "interface" { return "Interface" }
    if yname == "maintenance-point-type" { return "MaintenancePointType" }
    if yname == "mep-id" { return "MepId" }
    return ""
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetSegmentPath() string {
    return "maintenance-point"
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-name"] = maintenancePoint.DomainName
    leafs["level"] = maintenancePoint.Level
    leafs["service-name"] = maintenancePoint.ServiceName
    leafs["interface"] = maintenancePoint.Interface
    leafs["maintenance-point-type"] = maintenancePoint.MaintenancePointType
    leafs["mep-id"] = maintenancePoint.MepId
    return leafs
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetBundleName() string { return "cisco_ios_xr" }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetYangName() string { return "maintenance-point" }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) SetParent(parent types.Entity) { maintenancePoint.parent = parent }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetParent() types.Entity { return maintenancePoint.parent }

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint) GetParentYangName() string { return "maintenance-point" }

// Cfm_Global_GlobalConfigurationErrors
// Global configuration errors table
type Cfm_Global_GlobalConfigurationErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular configuration error. The type is slice of
    // Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError.
    GlobalConfigurationError []Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError
}

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetFilter() yfilter.YFilter { return globalConfigurationErrors.YFilter }

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) SetFilter(yf yfilter.YFilter) { globalConfigurationErrors.YFilter = yf }

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetGoName(yname string) string {
    if yname == "global-configuration-error" { return "GlobalConfigurationError" }
    return ""
}

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetSegmentPath() string {
    return "global-configuration-errors"
}

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-configuration-error" {
        for _, c := range globalConfigurationErrors.GlobalConfigurationError {
            if globalConfigurationErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError{}
        globalConfigurationErrors.GlobalConfigurationError = append(globalConfigurationErrors.GlobalConfigurationError, child)
        return &globalConfigurationErrors.GlobalConfigurationError[len(globalConfigurationErrors.GlobalConfigurationError)-1]
    }
    return nil
}

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range globalConfigurationErrors.GlobalConfigurationError {
        children[globalConfigurationErrors.GlobalConfigurationError[i].GetSegmentPath()] = &globalConfigurationErrors.GlobalConfigurationError[i]
    }
    return children
}

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetBundleName() string { return "cisco_ios_xr" }

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetYangName() string { return "global-configuration-errors" }

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) SetParent(parent types.Entity) { globalConfigurationErrors.parent = parent }

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetParent() types.Entity { return globalConfigurationErrors.parent }

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetParentYangName() string { return "global" }

// Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError
// Information about a particular configuration
// error
type Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // Level. The type is CfmBagMdLevel.
    Level interface{}

    // Service name. The type is string.
    ServiceName interface{}

    // The BD/XC is configured globally. The type is bool.
    BridgeDomainIsConfigured interface{}

    // The BD/XC could not be downloaded to L2FIB. The type is bool.
    L2FibDownloadError interface{}

    // BD/XC ID, or Service name if the Service is 'down-only'.
    BridgeDomainId Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId
}

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetFilter() yfilter.YFilter { return globalConfigurationError.YFilter }

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) SetFilter(yf yfilter.YFilter) { globalConfigurationError.YFilter = yf }

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "service" { return "Service" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "level" { return "Level" }
    if yname == "service-name" { return "ServiceName" }
    if yname == "bridge-domain-is-configured" { return "BridgeDomainIsConfigured" }
    if yname == "l2-fib-download-error" { return "L2FibDownloadError" }
    if yname == "bridge-domain-id" { return "BridgeDomainId" }
    return ""
}

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetSegmentPath() string {
    return "global-configuration-error" + "[domain='" + fmt.Sprintf("%v", globalConfigurationError.Domain) + "']" + "[service='" + fmt.Sprintf("%v", globalConfigurationError.Service) + "']"
}

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bridge-domain-id" {
        return &globalConfigurationError.BridgeDomainId
    }
    return nil
}

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bridge-domain-id"] = &globalConfigurationError.BridgeDomainId
    return children
}

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = globalConfigurationError.Domain
    leafs["service"] = globalConfigurationError.Service
    leafs["domain-name"] = globalConfigurationError.DomainName
    leafs["level"] = globalConfigurationError.Level
    leafs["service-name"] = globalConfigurationError.ServiceName
    leafs["bridge-domain-is-configured"] = globalConfigurationError.BridgeDomainIsConfigured
    leafs["l2-fib-download-error"] = globalConfigurationError.L2FibDownloadError
    return leafs
}

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetBundleName() string { return "cisco_ios_xr" }

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetYangName() string { return "global-configuration-error" }

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) SetParent(parent types.Entity) { globalConfigurationError.parent = parent }

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetParent() types.Entity { return globalConfigurationError.parent }

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetParentYangName() string { return "global-configuration-errors" }

// Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId
// BD/XC ID, or Service name if the Service is
// 'down-only'
type Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bridge domain identifier format. The type is CfmBagBdidFmt.
    BridgeDomainIdFormat interface{}

    // Name of the Bridge/XConnect Group. The type is string.
    Group interface{}

    // Name of the Bridge Domain/XConnect. The type is string.
    Name interface{}

    // Local Customer Edge Identifier (CE-ID). The type is interface{} with range:
    // 0..65535.
    CeId interface{}

    // Remote Customer Edge Identifier (CE-ID). The type is interface{} with
    // range: 0..65535.
    RemoteCeId interface{}

    // EVPN ID for VLAN-aware flexible cross-connects. The type is interface{}
    // with range: 0..4294967295.
    Evi interface{}
}

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetFilter() yfilter.YFilter { return bridgeDomainId.YFilter }

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) SetFilter(yf yfilter.YFilter) { bridgeDomainId.YFilter = yf }

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetGoName(yname string) string {
    if yname == "bridge-domain-id-format" { return "BridgeDomainIdFormat" }
    if yname == "group" { return "Group" }
    if yname == "name" { return "Name" }
    if yname == "ce-id" { return "CeId" }
    if yname == "remote-ce-id" { return "RemoteCeId" }
    if yname == "evi" { return "Evi" }
    return ""
}

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetSegmentPath() string {
    return "bridge-domain-id"
}

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bridge-domain-id-format"] = bridgeDomainId.BridgeDomainIdFormat
    leafs["group"] = bridgeDomainId.Group
    leafs["name"] = bridgeDomainId.Name
    leafs["ce-id"] = bridgeDomainId.CeId
    leafs["remote-ce-id"] = bridgeDomainId.RemoteCeId
    leafs["evi"] = bridgeDomainId.Evi
    return leafs
}

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetBundleName() string { return "cisco_ios_xr" }

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetYangName() string { return "bridge-domain-id" }

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) SetParent(parent types.Entity) { bridgeDomainId.parent = parent }

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetParent() types.Entity { return bridgeDomainId.parent }

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetParentYangName() string { return "global-configuration-error" }

// Cfm_Global_MepConfigurationErrors
// MEP configuration errors table
type Cfm_Global_MepConfigurationErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular configuration error. The type is slice of
    // Cfm_Global_MepConfigurationErrors_MepConfigurationError.
    MepConfigurationError []Cfm_Global_MepConfigurationErrors_MepConfigurationError
}

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetFilter() yfilter.YFilter { return mepConfigurationErrors.YFilter }

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) SetFilter(yf yfilter.YFilter) { mepConfigurationErrors.YFilter = yf }

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetGoName(yname string) string {
    if yname == "mep-configuration-error" { return "MepConfigurationError" }
    return ""
}

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetSegmentPath() string {
    return "mep-configuration-errors"
}

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mep-configuration-error" {
        for _, c := range mepConfigurationErrors.MepConfigurationError {
            if mepConfigurationErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_MepConfigurationErrors_MepConfigurationError{}
        mepConfigurationErrors.MepConfigurationError = append(mepConfigurationErrors.MepConfigurationError, child)
        return &mepConfigurationErrors.MepConfigurationError[len(mepConfigurationErrors.MepConfigurationError)-1]
    }
    return nil
}

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mepConfigurationErrors.MepConfigurationError {
        children[mepConfigurationErrors.MepConfigurationError[i].GetSegmentPath()] = &mepConfigurationErrors.MepConfigurationError[i]
    }
    return children
}

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetBundleName() string { return "cisco_ios_xr" }

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetYangName() string { return "mep-configuration-errors" }

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) SetParent(parent types.Entity) { mepConfigurationErrors.parent = parent }

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetParent() types.Entity { return mepConfigurationErrors.parent }

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetParentYangName() string { return "global" }

// Cfm_Global_MepConfigurationErrors_MepConfigurationError
// Information about a particular configuration
// error
type Cfm_Global_MepConfigurationErrors_MepConfigurationError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}

    // Interval between CCMs sent on this MEP. The type is CfmBagCcmInterval.
    CcmInterval interface{}

    // The MEP's Domain is not configured. The type is bool.
    NoDomain interface{}

    // The MEP's Service is not configured. The type is bool.
    NoService interface{}

    // The MEP's EFP is not in the Service's Bridge Domain. The type is bool.
    BridgeDomainMismatch interface{}

    // Another MEP facing in the same direction is at the same Maintenance Level.
    // The type is bool.
    LevelConflict interface{}

    // CCM Interval is less than minimum interval supported by hardware. The type
    // is bool.
    CcmIntervalNotSupported interface{}

    // Offload resource limits have been exceeded. The type is bool.
    OffloadOutOfResources interface{}

    // Multiple offloaded MEPs on the same interface are not supported. The type
    // is bool.
    OffloadMultipleLocalMep interface{}

    // The MEP should be offloaded but crosscheck has not been configured. The
    // type is bool.
    OffloadNoCrossCheck interface{}

    // The MEP should be offloaded but multiple crosscheck MEPs have been
    // configured, and this is not supported. The type is bool.
    OffloadMultiplePeerMeps interface{}

    // The MEP direction does not support offload. The type is bool.
    OffloadMepDirectionNotSupported interface{}

    // AIS is configured on the same interface as the down MEP. The type is bool.
    AisConfigured interface{}

    // The MEP is configured in a domain at level 0, on a bundle interface or
    // sub-interface.  This is not supported. The type is bool.
    BundleLevel0 interface{}

    // A BD/XC specified in the MEG config, but it does not exist globally. The
    // type is bool.
    BridgeDomainNotInBdInfra interface{}

    // The configured MAID format is not supported for hardware offload. The type
    // is bool.
    MaidFormatNotSupported interface{}

    // The platform returned a fatal error when passed the offload session. The
    // type is bool.
    FatalOffloadError interface{}

    // A satellite limitation is preventing MEP being offloaded to satellite. The
    // type is bool.
    SatelliteLimitation interface{}

    // In-progress Ethernet SLA loopback operations are disabled due to satellite
    // having loopback responder-only capabilities. The type is bool.
    SlaLoopbackOperationsDisabled interface{}

    // In-progress Ethernet SLA synthetic loss measurement operations are disabled
    // due to satellite having synthetic loss measurement responder-only
    // capabilities. The type is bool.
    SlaSyntheticLossOperationsDisabled interface{}

    // In-progress Ethernet SLA delay measurement operations are disabled due to
    // satellite having delay measurement responder-only capabilities. The type is
    // bool.
    SlaDelayMeasurementOperationsDisabled interface{}

    // The EFP doesn't have a valid MAC address yet. This will also get set if the
    // MAC address we have is a multicast address. The type is bool.
    NoValidMacAddress interface{}

    // We haven't yet been able to look up the interface type to find whether the
    // interface is a bundle. The type is bool.
    NoInterfaceType interface{}

    // The EFP has been deleted from IM. The type is bool.
    NotInIm interface{}

    // The EFP is a bundle and the mLACP mode is not yet known. The type is bool.
    NoMlacp interface{}

    // Error string returned from satellite. The type is string.
    SatelliteErrorString interface{}

    // ID of the satellite. The type is interface{} with range: 0..65535.
    SatelliteId interface{}

    // The MEP that has errors.
    Mep Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep

    // BD/XC ID for the MEP's Service, or Service name if the Service is
    // 'down-only'.
    ServiceBridgeDomain Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain

    // ID of the BD/XC that the MEP's EFP is in, if any.
    InterfaceBridgeDomain Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain

    // Satellite Capabilities.
    SatelliteCapabilities Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities
}

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetFilter() yfilter.YFilter { return mepConfigurationError.YFilter }

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) SetFilter(yf yfilter.YFilter) { mepConfigurationError.YFilter = yf }

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "service" { return "Service" }
    if yname == "interface" { return "Interface" }
    if yname == "ccm-interval" { return "CcmInterval" }
    if yname == "no-domain" { return "NoDomain" }
    if yname == "no-service" { return "NoService" }
    if yname == "bridge-domain-mismatch" { return "BridgeDomainMismatch" }
    if yname == "level-conflict" { return "LevelConflict" }
    if yname == "ccm-interval-not-supported" { return "CcmIntervalNotSupported" }
    if yname == "offload-out-of-resources" { return "OffloadOutOfResources" }
    if yname == "offload-multiple-local-mep" { return "OffloadMultipleLocalMep" }
    if yname == "offload-no-cross-check" { return "OffloadNoCrossCheck" }
    if yname == "offload-multiple-peer-meps" { return "OffloadMultiplePeerMeps" }
    if yname == "offload-mep-direction-not-supported" { return "OffloadMepDirectionNotSupported" }
    if yname == "ais-configured" { return "AisConfigured" }
    if yname == "bundle-level0" { return "BundleLevel0" }
    if yname == "bridge-domain-not-in-bd-infra" { return "BridgeDomainNotInBdInfra" }
    if yname == "maid-format-not-supported" { return "MaidFormatNotSupported" }
    if yname == "fatal-offload-error" { return "FatalOffloadError" }
    if yname == "satellite-limitation" { return "SatelliteLimitation" }
    if yname == "sla-loopback-operations-disabled" { return "SlaLoopbackOperationsDisabled" }
    if yname == "sla-synthetic-loss-operations-disabled" { return "SlaSyntheticLossOperationsDisabled" }
    if yname == "sla-delay-measurement-operations-disabled" { return "SlaDelayMeasurementOperationsDisabled" }
    if yname == "no-valid-mac-address" { return "NoValidMacAddress" }
    if yname == "no-interface-type" { return "NoInterfaceType" }
    if yname == "not-in-im" { return "NotInIm" }
    if yname == "no-mlacp" { return "NoMlacp" }
    if yname == "satellite-error-string" { return "SatelliteErrorString" }
    if yname == "satellite-id" { return "SatelliteId" }
    if yname == "mep" { return "Mep" }
    if yname == "service-bridge-domain" { return "ServiceBridgeDomain" }
    if yname == "interface-bridge-domain" { return "InterfaceBridgeDomain" }
    if yname == "satellite-capabilities" { return "SatelliteCapabilities" }
    return ""
}

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetSegmentPath() string {
    return "mep-configuration-error" + "[domain='" + fmt.Sprintf("%v", mepConfigurationError.Domain) + "']" + "[service='" + fmt.Sprintf("%v", mepConfigurationError.Service) + "']" + "[interface='" + fmt.Sprintf("%v", mepConfigurationError.Interface) + "']"
}

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mep" {
        return &mepConfigurationError.Mep
    }
    if childYangName == "service-bridge-domain" {
        return &mepConfigurationError.ServiceBridgeDomain
    }
    if childYangName == "interface-bridge-domain" {
        return &mepConfigurationError.InterfaceBridgeDomain
    }
    if childYangName == "satellite-capabilities" {
        return &mepConfigurationError.SatelliteCapabilities
    }
    return nil
}

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mep"] = &mepConfigurationError.Mep
    children["service-bridge-domain"] = &mepConfigurationError.ServiceBridgeDomain
    children["interface-bridge-domain"] = &mepConfigurationError.InterfaceBridgeDomain
    children["satellite-capabilities"] = &mepConfigurationError.SatelliteCapabilities
    return children
}

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = mepConfigurationError.Domain
    leafs["service"] = mepConfigurationError.Service
    leafs["interface"] = mepConfigurationError.Interface
    leafs["ccm-interval"] = mepConfigurationError.CcmInterval
    leafs["no-domain"] = mepConfigurationError.NoDomain
    leafs["no-service"] = mepConfigurationError.NoService
    leafs["bridge-domain-mismatch"] = mepConfigurationError.BridgeDomainMismatch
    leafs["level-conflict"] = mepConfigurationError.LevelConflict
    leafs["ccm-interval-not-supported"] = mepConfigurationError.CcmIntervalNotSupported
    leafs["offload-out-of-resources"] = mepConfigurationError.OffloadOutOfResources
    leafs["offload-multiple-local-mep"] = mepConfigurationError.OffloadMultipleLocalMep
    leafs["offload-no-cross-check"] = mepConfigurationError.OffloadNoCrossCheck
    leafs["offload-multiple-peer-meps"] = mepConfigurationError.OffloadMultiplePeerMeps
    leafs["offload-mep-direction-not-supported"] = mepConfigurationError.OffloadMepDirectionNotSupported
    leafs["ais-configured"] = mepConfigurationError.AisConfigured
    leafs["bundle-level0"] = mepConfigurationError.BundleLevel0
    leafs["bridge-domain-not-in-bd-infra"] = mepConfigurationError.BridgeDomainNotInBdInfra
    leafs["maid-format-not-supported"] = mepConfigurationError.MaidFormatNotSupported
    leafs["fatal-offload-error"] = mepConfigurationError.FatalOffloadError
    leafs["satellite-limitation"] = mepConfigurationError.SatelliteLimitation
    leafs["sla-loopback-operations-disabled"] = mepConfigurationError.SlaLoopbackOperationsDisabled
    leafs["sla-synthetic-loss-operations-disabled"] = mepConfigurationError.SlaSyntheticLossOperationsDisabled
    leafs["sla-delay-measurement-operations-disabled"] = mepConfigurationError.SlaDelayMeasurementOperationsDisabled
    leafs["no-valid-mac-address"] = mepConfigurationError.NoValidMacAddress
    leafs["no-interface-type"] = mepConfigurationError.NoInterfaceType
    leafs["not-in-im"] = mepConfigurationError.NotInIm
    leafs["no-mlacp"] = mepConfigurationError.NoMlacp
    leafs["satellite-error-string"] = mepConfigurationError.SatelliteErrorString
    leafs["satellite-id"] = mepConfigurationError.SatelliteId
    return leafs
}

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetBundleName() string { return "cisco_ios_xr" }

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetYangName() string { return "mep-configuration-error" }

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) SetParent(parent types.Entity) { mepConfigurationError.parent = parent }

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetParent() types.Entity { return mepConfigurationError.parent }

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetParentYangName() string { return "mep-configuration-errors" }

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep
// The MEP that has errors
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain name. The type is string.
    DomainName interface{}

    // Domain level. The type is CfmBagMdLevel.
    Level interface{}

    // Service name. The type is string.
    ServiceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Type of Maintenance Point. The type is CfmMaMpVariety.
    MaintenancePointType interface{}

    // MEP ID. The type is interface{} with range: 0..65535.
    MepId interface{}
}

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetFilter() yfilter.YFilter { return mep.YFilter }

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) SetFilter(yf yfilter.YFilter) { mep.YFilter = yf }

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetGoName(yname string) string {
    if yname == "domain-name" { return "DomainName" }
    if yname == "level" { return "Level" }
    if yname == "service-name" { return "ServiceName" }
    if yname == "interface" { return "Interface" }
    if yname == "maintenance-point-type" { return "MaintenancePointType" }
    if yname == "mep-id" { return "MepId" }
    return ""
}

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetSegmentPath() string {
    return "mep"
}

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-name"] = mep.DomainName
    leafs["level"] = mep.Level
    leafs["service-name"] = mep.ServiceName
    leafs["interface"] = mep.Interface
    leafs["maintenance-point-type"] = mep.MaintenancePointType
    leafs["mep-id"] = mep.MepId
    return leafs
}

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetBundleName() string { return "cisco_ios_xr" }

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetYangName() string { return "mep" }

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) SetParent(parent types.Entity) { mep.parent = parent }

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetParent() types.Entity { return mep.parent }

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetParentYangName() string { return "mep-configuration-error" }

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain
// BD/XC ID for the MEP's Service, or Service name
// if the Service is 'down-only'
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bridge domain identifier format. The type is CfmBagBdidFmt.
    BridgeDomainIdFormat interface{}

    // Name of the Bridge/XConnect Group. The type is string.
    Group interface{}

    // Name of the Bridge Domain/XConnect. The type is string.
    Name interface{}

    // Local Customer Edge Identifier (CE-ID). The type is interface{} with range:
    // 0..65535.
    CeId interface{}

    // Remote Customer Edge Identifier (CE-ID). The type is interface{} with
    // range: 0..65535.
    RemoteCeId interface{}

    // EVPN ID for VLAN-aware flexible cross-connects. The type is interface{}
    // with range: 0..4294967295.
    Evi interface{}
}

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetFilter() yfilter.YFilter { return serviceBridgeDomain.YFilter }

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) SetFilter(yf yfilter.YFilter) { serviceBridgeDomain.YFilter = yf }

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetGoName(yname string) string {
    if yname == "bridge-domain-id-format" { return "BridgeDomainIdFormat" }
    if yname == "group" { return "Group" }
    if yname == "name" { return "Name" }
    if yname == "ce-id" { return "CeId" }
    if yname == "remote-ce-id" { return "RemoteCeId" }
    if yname == "evi" { return "Evi" }
    return ""
}

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetSegmentPath() string {
    return "service-bridge-domain"
}

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bridge-domain-id-format"] = serviceBridgeDomain.BridgeDomainIdFormat
    leafs["group"] = serviceBridgeDomain.Group
    leafs["name"] = serviceBridgeDomain.Name
    leafs["ce-id"] = serviceBridgeDomain.CeId
    leafs["remote-ce-id"] = serviceBridgeDomain.RemoteCeId
    leafs["evi"] = serviceBridgeDomain.Evi
    return leafs
}

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetBundleName() string { return "cisco_ios_xr" }

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetYangName() string { return "service-bridge-domain" }

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) SetParent(parent types.Entity) { serviceBridgeDomain.parent = parent }

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetParent() types.Entity { return serviceBridgeDomain.parent }

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetParentYangName() string { return "mep-configuration-error" }

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain
// ID of the BD/XC that the MEP's EFP is in, if any
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bridge domain identifier format. The type is CfmBagBdidFmt.
    BridgeDomainIdFormat interface{}

    // Name of the Bridge/XConnect Group. The type is string.
    Group interface{}

    // Name of the Bridge Domain/XConnect. The type is string.
    Name interface{}

    // Local Customer Edge Identifier (CE-ID). The type is interface{} with range:
    // 0..65535.
    CeId interface{}

    // Remote Customer Edge Identifier (CE-ID). The type is interface{} with
    // range: 0..65535.
    RemoteCeId interface{}

    // EVPN ID for VLAN-aware flexible cross-connects. The type is interface{}
    // with range: 0..4294967295.
    Evi interface{}
}

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetFilter() yfilter.YFilter { return interfaceBridgeDomain.YFilter }

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) SetFilter(yf yfilter.YFilter) { interfaceBridgeDomain.YFilter = yf }

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetGoName(yname string) string {
    if yname == "bridge-domain-id-format" { return "BridgeDomainIdFormat" }
    if yname == "group" { return "Group" }
    if yname == "name" { return "Name" }
    if yname == "ce-id" { return "CeId" }
    if yname == "remote-ce-id" { return "RemoteCeId" }
    if yname == "evi" { return "Evi" }
    return ""
}

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetSegmentPath() string {
    return "interface-bridge-domain"
}

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bridge-domain-id-format"] = interfaceBridgeDomain.BridgeDomainIdFormat
    leafs["group"] = interfaceBridgeDomain.Group
    leafs["name"] = interfaceBridgeDomain.Name
    leafs["ce-id"] = interfaceBridgeDomain.CeId
    leafs["remote-ce-id"] = interfaceBridgeDomain.RemoteCeId
    leafs["evi"] = interfaceBridgeDomain.Evi
    return leafs
}

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetYangName() string { return "interface-bridge-domain" }

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) SetParent(parent types.Entity) { interfaceBridgeDomain.parent = parent }

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetParent() types.Entity { return interfaceBridgeDomain.parent }

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetParentYangName() string { return "mep-configuration-error" }

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities
// Satellite Capabilities
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Loopback.
    Loopback Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback

    // Delay Measurement.
    DelayMeasurement Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement

    // Synthetic Loss Measurement.
    SyntheticLossMeasurement Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement
}

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetFilter() yfilter.YFilter { return satelliteCapabilities.YFilter }

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) SetFilter(yf yfilter.YFilter) { satelliteCapabilities.YFilter = yf }

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetGoName(yname string) string {
    if yname == "loopback" { return "Loopback" }
    if yname == "delay-measurement" { return "DelayMeasurement" }
    if yname == "synthetic-loss-measurement" { return "SyntheticLossMeasurement" }
    return ""
}

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetSegmentPath() string {
    return "satellite-capabilities"
}

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "loopback" {
        return &satelliteCapabilities.Loopback
    }
    if childYangName == "delay-measurement" {
        return &satelliteCapabilities.DelayMeasurement
    }
    if childYangName == "synthetic-loss-measurement" {
        return &satelliteCapabilities.SyntheticLossMeasurement
    }
    return nil
}

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["loopback"] = &satelliteCapabilities.Loopback
    children["delay-measurement"] = &satelliteCapabilities.DelayMeasurement
    children["synthetic-loss-measurement"] = &satelliteCapabilities.SyntheticLossMeasurement
    return children
}

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetBundleName() string { return "cisco_ios_xr" }

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetYangName() string { return "satellite-capabilities" }

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) SetParent(parent types.Entity) { satelliteCapabilities.parent = parent }

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetParent() types.Entity { return satelliteCapabilities.parent }

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetParentYangName() string { return "mep-configuration-error" }

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback
// Loopback
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Responder. The type is bool.
    Responder interface{}

    // Controller. The type is bool.
    Controller interface{}
}

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetFilter() yfilter.YFilter { return loopback.YFilter }

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) SetFilter(yf yfilter.YFilter) { loopback.YFilter = yf }

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetGoName(yname string) string {
    if yname == "responder" { return "Responder" }
    if yname == "controller" { return "Controller" }
    return ""
}

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetSegmentPath() string {
    return "loopback"
}

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["responder"] = loopback.Responder
    leafs["controller"] = loopback.Controller
    return leafs
}

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetBundleName() string { return "cisco_ios_xr" }

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetYangName() string { return "loopback" }

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) SetParent(parent types.Entity) { loopback.parent = parent }

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetParent() types.Entity { return loopback.parent }

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetParentYangName() string { return "satellite-capabilities" }

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement
// Delay Measurement
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Responder. The type is bool.
    Responder interface{}

    // Controller. The type is bool.
    Controller interface{}
}

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetFilter() yfilter.YFilter { return delayMeasurement.YFilter }

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) SetFilter(yf yfilter.YFilter) { delayMeasurement.YFilter = yf }

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetGoName(yname string) string {
    if yname == "responder" { return "Responder" }
    if yname == "controller" { return "Controller" }
    return ""
}

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetSegmentPath() string {
    return "delay-measurement"
}

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["responder"] = delayMeasurement.Responder
    leafs["controller"] = delayMeasurement.Controller
    return leafs
}

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetBundleName() string { return "cisco_ios_xr" }

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetYangName() string { return "delay-measurement" }

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) SetParent(parent types.Entity) { delayMeasurement.parent = parent }

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetParent() types.Entity { return delayMeasurement.parent }

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetParentYangName() string { return "satellite-capabilities" }

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement
// Synthetic Loss Measurement
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Responder. The type is bool.
    Responder interface{}

    // Controller. The type is bool.
    Controller interface{}
}

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetFilter() yfilter.YFilter { return syntheticLossMeasurement.YFilter }

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) SetFilter(yf yfilter.YFilter) { syntheticLossMeasurement.YFilter = yf }

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetGoName(yname string) string {
    if yname == "responder" { return "Responder" }
    if yname == "controller" { return "Controller" }
    return ""
}

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetSegmentPath() string {
    return "synthetic-loss-measurement"
}

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["responder"] = syntheticLossMeasurement.Responder
    leafs["controller"] = syntheticLossMeasurement.Controller
    return leafs
}

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetBundleName() string { return "cisco_ios_xr" }

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetYangName() string { return "synthetic-loss-measurement" }

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) SetParent(parent types.Entity) { syntheticLossMeasurement.parent = parent }

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetParent() types.Entity { return syntheticLossMeasurement.parent }

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetParentYangName() string { return "satellite-capabilities" }

// Cfm_Global_TracerouteCaches
// Traceroute Cache table
type Cfm_Global_TracerouteCaches struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular traceroute operation. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache.
    TracerouteCache []Cfm_Global_TracerouteCaches_TracerouteCache
}

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetFilter() yfilter.YFilter { return tracerouteCaches.YFilter }

func (tracerouteCaches *Cfm_Global_TracerouteCaches) SetFilter(yf yfilter.YFilter) { tracerouteCaches.YFilter = yf }

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetGoName(yname string) string {
    if yname == "traceroute-cache" { return "TracerouteCache" }
    return ""
}

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetSegmentPath() string {
    return "traceroute-caches"
}

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traceroute-cache" {
        for _, c := range tracerouteCaches.TracerouteCache {
            if tracerouteCaches.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_TracerouteCaches_TracerouteCache{}
        tracerouteCaches.TracerouteCache = append(tracerouteCaches.TracerouteCache, child)
        return &tracerouteCaches.TracerouteCache[len(tracerouteCaches.TracerouteCache)-1]
    }
    return nil
}

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tracerouteCaches.TracerouteCache {
        children[tracerouteCaches.TracerouteCache[i].GetSegmentPath()] = &tracerouteCaches.TracerouteCache[i]
    }
    return children
}

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetBundleName() string { return "cisco_ios_xr" }

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetYangName() string { return "traceroute-caches" }

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tracerouteCaches *Cfm_Global_TracerouteCaches) SetParent(parent types.Entity) { tracerouteCaches.parent = parent }

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetParent() types.Entity { return tracerouteCaches.parent }

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetParentYangName() string { return "global" }

// Cfm_Global_TracerouteCaches_TracerouteCache
// Information about a particular traceroute
// operation
type Cfm_Global_TracerouteCaches_TracerouteCache struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. MEP ID. The type is interface{} with range:
    // 1..8191.
    MepId interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}

    // This attribute is a key. Transaction ID. The type is interface{} with
    // range: -2147483648..2147483647.
    TransactionId interface{}

    // Count of ignored replies for this request. The type is interface{} with
    // range: 0..4294967295.
    RepliesDropped interface{}

    // Information about the traceroute operation.
    TracerouteInformation Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation

    // Received linktrace replies. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply.
    LinktraceReply []Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply

    // Received exploratory linktrace replies. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply.
    ExploratoryLinktraceReply []Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply
}

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetFilter() yfilter.YFilter { return tracerouteCache.YFilter }

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) SetFilter(yf yfilter.YFilter) { tracerouteCache.YFilter = yf }

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "service" { return "Service" }
    if yname == "mep-id" { return "MepId" }
    if yname == "interface" { return "Interface" }
    if yname == "transaction-id" { return "TransactionId" }
    if yname == "replies-dropped" { return "RepliesDropped" }
    if yname == "traceroute-information" { return "TracerouteInformation" }
    if yname == "linktrace-reply" { return "LinktraceReply" }
    if yname == "exploratory-linktrace-reply" { return "ExploratoryLinktraceReply" }
    return ""
}

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetSegmentPath() string {
    return "traceroute-cache" + "[domain='" + fmt.Sprintf("%v", tracerouteCache.Domain) + "']" + "[service='" + fmt.Sprintf("%v", tracerouteCache.Service) + "']" + "[mep-id='" + fmt.Sprintf("%v", tracerouteCache.MepId) + "']" + "[interface='" + fmt.Sprintf("%v", tracerouteCache.Interface) + "']" + "[transaction-id='" + fmt.Sprintf("%v", tracerouteCache.TransactionId) + "']"
}

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traceroute-information" {
        return &tracerouteCache.TracerouteInformation
    }
    if childYangName == "linktrace-reply" {
        for _, c := range tracerouteCache.LinktraceReply {
            if tracerouteCache.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply{}
        tracerouteCache.LinktraceReply = append(tracerouteCache.LinktraceReply, child)
        return &tracerouteCache.LinktraceReply[len(tracerouteCache.LinktraceReply)-1]
    }
    if childYangName == "exploratory-linktrace-reply" {
        for _, c := range tracerouteCache.ExploratoryLinktraceReply {
            if tracerouteCache.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply{}
        tracerouteCache.ExploratoryLinktraceReply = append(tracerouteCache.ExploratoryLinktraceReply, child)
        return &tracerouteCache.ExploratoryLinktraceReply[len(tracerouteCache.ExploratoryLinktraceReply)-1]
    }
    return nil
}

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traceroute-information"] = &tracerouteCache.TracerouteInformation
    for i := range tracerouteCache.LinktraceReply {
        children[tracerouteCache.LinktraceReply[i].GetSegmentPath()] = &tracerouteCache.LinktraceReply[i]
    }
    for i := range tracerouteCache.ExploratoryLinktraceReply {
        children[tracerouteCache.ExploratoryLinktraceReply[i].GetSegmentPath()] = &tracerouteCache.ExploratoryLinktraceReply[i]
    }
    return children
}

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = tracerouteCache.Domain
    leafs["service"] = tracerouteCache.Service
    leafs["mep-id"] = tracerouteCache.MepId
    leafs["interface"] = tracerouteCache.Interface
    leafs["transaction-id"] = tracerouteCache.TransactionId
    leafs["replies-dropped"] = tracerouteCache.RepliesDropped
    return leafs
}

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetBundleName() string { return "cisco_ios_xr" }

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetYangName() string { return "traceroute-cache" }

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) SetParent(parent types.Entity) { tracerouteCache.parent = parent }

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetParent() types.Entity { return tracerouteCache.parent }

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetParentYangName() string { return "traceroute-caches" }

// Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation
// Information about the traceroute operation
type Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maintenance domain name. The type is string.
    Domain interface{}

    // Service name. The type is string.
    Service interface{}

    // Maintenance level. The type is CfmBagMdLevel.
    Level interface{}

    // Source MEP ID. The type is interface{} with range: 0..65535.
    SourceMepId interface{}

    // Source interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Source MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SourceMacAddress interface{}

    // Target MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    TargetMacAddress interface{}

    // Directed MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DirectedMacAddress interface{}

    // Target MEP ID. The type is interface{} with range: 0..65535.
    TargetMepId interface{}

    // Timestamp of initiation time (seconds). The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    Timestamp interface{}

    // Time to live. The type is interface{} with range: 0..255.
    Ttl interface{}

    // Transaction ID. The type is interface{} with range: 0..4294967295.
    TransactionId interface{}

    // Options affecting traceroute behavior.
    Options Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options
}

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetFilter() yfilter.YFilter { return tracerouteInformation.YFilter }

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) SetFilter(yf yfilter.YFilter) { tracerouteInformation.YFilter = yf }

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "service" { return "Service" }
    if yname == "level" { return "Level" }
    if yname == "source-mep-id" { return "SourceMepId" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "source-mac-address" { return "SourceMacAddress" }
    if yname == "target-mac-address" { return "TargetMacAddress" }
    if yname == "directed-mac-address" { return "DirectedMacAddress" }
    if yname == "target-mep-id" { return "TargetMepId" }
    if yname == "timestamp" { return "Timestamp" }
    if yname == "ttl" { return "Ttl" }
    if yname == "transaction-id" { return "TransactionId" }
    if yname == "options" { return "Options" }
    return ""
}

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetSegmentPath() string {
    return "traceroute-information"
}

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "options" {
        return &tracerouteInformation.Options
    }
    return nil
}

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["options"] = &tracerouteInformation.Options
    return children
}

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = tracerouteInformation.Domain
    leafs["service"] = tracerouteInformation.Service
    leafs["level"] = tracerouteInformation.Level
    leafs["source-mep-id"] = tracerouteInformation.SourceMepId
    leafs["source-interface"] = tracerouteInformation.SourceInterface
    leafs["source-mac-address"] = tracerouteInformation.SourceMacAddress
    leafs["target-mac-address"] = tracerouteInformation.TargetMacAddress
    leafs["directed-mac-address"] = tracerouteInformation.DirectedMacAddress
    leafs["target-mep-id"] = tracerouteInformation.TargetMepId
    leafs["timestamp"] = tracerouteInformation.Timestamp
    leafs["ttl"] = tracerouteInformation.Ttl
    leafs["transaction-id"] = tracerouteInformation.TransactionId
    return leafs
}

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetBundleName() string { return "cisco_ios_xr" }

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetYangName() string { return "traceroute-information" }

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) SetParent(parent types.Entity) { tracerouteInformation.parent = parent }

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetParent() types.Entity { return tracerouteInformation.parent }

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetParentYangName() string { return "traceroute-cache" }

// Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options
// Options affecting traceroute behavior
type Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mode. The type is CfmPmLtMode.
    Mode interface{}

    // Options for a basic IEEE 802.1ag Linktrace.
    BasicOptions Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions

    // Options for an Exploratory Linktrace.
    ExploratoryOptions Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions
}

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetFilter() yfilter.YFilter { return options.YFilter }

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) SetFilter(yf yfilter.YFilter) { options.YFilter = yf }

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    if yname == "basic-options" { return "BasicOptions" }
    if yname == "exploratory-options" { return "ExploratoryOptions" }
    return ""
}

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetSegmentPath() string {
    return "options"
}

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-options" {
        return &options.BasicOptions
    }
    if childYangName == "exploratory-options" {
        return &options.ExploratoryOptions
    }
    return nil
}

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-options"] = &options.BasicOptions
    children["exploratory-options"] = &options.ExploratoryOptions
    return children
}

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode"] = options.Mode
    return leafs
}

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetBundleName() string { return "cisco_ios_xr" }

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetYangName() string { return "options" }

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) SetParent(parent types.Entity) { options.parent = parent }

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetParent() types.Entity { return options.parent }

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetParentYangName() string { return "traceroute-information" }

// Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions
// Options for a basic IEEE 802.1ag Linktrace
type Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Traceroute was initiated automatically. The type is bool.
    IsAuto interface{}

    // Only use the Filtering Database for forwarding lookups. The type is bool.
    FdbOnly interface{}
}

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetFilter() yfilter.YFilter { return basicOptions.YFilter }

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) SetFilter(yf yfilter.YFilter) { basicOptions.YFilter = yf }

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetGoName(yname string) string {
    if yname == "is-auto" { return "IsAuto" }
    if yname == "fdb-only" { return "FdbOnly" }
    return ""
}

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetSegmentPath() string {
    return "basic-options"
}

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-auto"] = basicOptions.IsAuto
    leafs["fdb-only"] = basicOptions.FdbOnly
    return leafs
}

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetBundleName() string { return "cisco_ios_xr" }

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetYangName() string { return "basic-options" }

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) SetParent(parent types.Entity) { basicOptions.parent = parent }

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetParent() types.Entity { return basicOptions.parent }

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetParentYangName() string { return "options" }

// Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions
// Options for an Exploratory Linktrace
type Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Delay model for delay calculations. The type is CfmPmEltDelayModel.
    DelayModel interface{}

    // Constant Factor for delay calculations. The type is interface{} with range:
    // 0..4294967295.
    DelayConstantFactor interface{}

    // Reply Filtering mode used by responders. The type is CfmPmElmReplyFilter.
    ReplyFilter interface{}
}

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetFilter() yfilter.YFilter { return exploratoryOptions.YFilter }

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) SetFilter(yf yfilter.YFilter) { exploratoryOptions.YFilter = yf }

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetGoName(yname string) string {
    if yname == "delay-model" { return "DelayModel" }
    if yname == "delay-constant-factor" { return "DelayConstantFactor" }
    if yname == "reply-filter" { return "ReplyFilter" }
    return ""
}

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetSegmentPath() string {
    return "exploratory-options"
}

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["delay-model"] = exploratoryOptions.DelayModel
    leafs["delay-constant-factor"] = exploratoryOptions.DelayConstantFactor
    leafs["reply-filter"] = exploratoryOptions.ReplyFilter
    return leafs
}

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetBundleName() string { return "cisco_ios_xr" }

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetYangName() string { return "exploratory-options" }

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) SetParent(parent types.Entity) { exploratoryOptions.parent = parent }

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetParent() types.Entity { return exploratoryOptions.parent }

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetParentYangName() string { return "options" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply
// Received linktrace replies
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Undecoded frame. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    RawData interface{}

    // Frame header.
    Header Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header

    // Sender ID TLV.
    SenderId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId

    // Egress ID TLV.
    EgressId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId

    // Reply ingress TLV.
    ReplyIngress Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress

    // Reply egress TLV.
    ReplyEgress Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress

    // Last hop ID.
    LastHop Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop

    // Organizational-specific TLVs. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv.
    OrganizationSpecificTlv []Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv

    // Unknown TLVs. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv.
    UnknownTlv []Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv
}

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetFilter() yfilter.YFilter { return linktraceReply.YFilter }

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) SetFilter(yf yfilter.YFilter) { linktraceReply.YFilter = yf }

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetGoName(yname string) string {
    if yname == "raw-data" { return "RawData" }
    if yname == "header" { return "Header" }
    if yname == "sender-id" { return "SenderId" }
    if yname == "egress-id" { return "EgressId" }
    if yname == "reply-ingress" { return "ReplyIngress" }
    if yname == "reply-egress" { return "ReplyEgress" }
    if yname == "last-hop" { return "LastHop" }
    if yname == "organization-specific-tlv" { return "OrganizationSpecificTlv" }
    if yname == "unknown-tlv" { return "UnknownTlv" }
    return ""
}

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetSegmentPath() string {
    return "linktrace-reply"
}

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &linktraceReply.Header
    }
    if childYangName == "sender-id" {
        return &linktraceReply.SenderId
    }
    if childYangName == "egress-id" {
        return &linktraceReply.EgressId
    }
    if childYangName == "reply-ingress" {
        return &linktraceReply.ReplyIngress
    }
    if childYangName == "reply-egress" {
        return &linktraceReply.ReplyEgress
    }
    if childYangName == "last-hop" {
        return &linktraceReply.LastHop
    }
    if childYangName == "organization-specific-tlv" {
        for _, c := range linktraceReply.OrganizationSpecificTlv {
            if linktraceReply.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv{}
        linktraceReply.OrganizationSpecificTlv = append(linktraceReply.OrganizationSpecificTlv, child)
        return &linktraceReply.OrganizationSpecificTlv[len(linktraceReply.OrganizationSpecificTlv)-1]
    }
    if childYangName == "unknown-tlv" {
        for _, c := range linktraceReply.UnknownTlv {
            if linktraceReply.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv{}
        linktraceReply.UnknownTlv = append(linktraceReply.UnknownTlv, child)
        return &linktraceReply.UnknownTlv[len(linktraceReply.UnknownTlv)-1]
    }
    return nil
}

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &linktraceReply.Header
    children["sender-id"] = &linktraceReply.SenderId
    children["egress-id"] = &linktraceReply.EgressId
    children["reply-ingress"] = &linktraceReply.ReplyIngress
    children["reply-egress"] = &linktraceReply.ReplyEgress
    children["last-hop"] = &linktraceReply.LastHop
    for i := range linktraceReply.OrganizationSpecificTlv {
        children[linktraceReply.OrganizationSpecificTlv[i].GetSegmentPath()] = &linktraceReply.OrganizationSpecificTlv[i]
    }
    for i := range linktraceReply.UnknownTlv {
        children[linktraceReply.UnknownTlv[i].GetSegmentPath()] = &linktraceReply.UnknownTlv[i]
    }
    return children
}

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["raw-data"] = linktraceReply.RawData
    return leafs
}

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetBundleName() string { return "cisco_ios_xr" }

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetYangName() string { return "linktrace-reply" }

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) SetParent(parent types.Entity) { linktraceReply.parent = parent }

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetParent() types.Entity { return linktraceReply.parent }

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetParentYangName() string { return "traceroute-cache" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header
// Frame header
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MD level. The type is CfmBagMdLevel.
    Level interface{}

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Use filtering DB only. The type is bool.
    UseFdbOnly interface{}

    // LTR was forwarded. The type is bool.
    Forwarded interface{}

    // Terminal MEP reached. The type is bool.
    TerminalMep interface{}

    // Transaction ID. The type is interface{} with range: 0..4294967295.
    TransactionId interface{}

    // TTL. The type is interface{} with range: 0..255.
    Ttl interface{}

    // Relay action. The type is CfmPmRelayAction.
    RelayAction interface{}
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "version" { return "Version" }
    if yname == "use-fdb-only" { return "UseFdbOnly" }
    if yname == "forwarded" { return "Forwarded" }
    if yname == "terminal-mep" { return "TerminalMep" }
    if yname == "transaction-id" { return "TransactionId" }
    if yname == "ttl" { return "Ttl" }
    if yname == "relay-action" { return "RelayAction" }
    return ""
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetSegmentPath() string {
    return "header"
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = header.Level
    leafs["version"] = header.Version
    leafs["use-fdb-only"] = header.UseFdbOnly
    leafs["forwarded"] = header.Forwarded
    leafs["terminal-mep"] = header.TerminalMep
    leafs["transaction-id"] = header.TransactionId
    leafs["ttl"] = header.Ttl
    leafs["relay-action"] = header.RelayAction
    return leafs
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetYangName() string { return "header" }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetParent() types.Entity { return header.parent }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetParentYangName() string { return "linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId
// Sender ID TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Management address domain. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ManagementAddressDomain interface{}

    // Management address. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ManagementAddress interface{}

    // Chassis ID.
    ChassisId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetFilter() yfilter.YFilter { return senderId.YFilter }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) SetFilter(yf yfilter.YFilter) { senderId.YFilter = yf }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetGoName(yname string) string {
    if yname == "management-address-domain" { return "ManagementAddressDomain" }
    if yname == "management-address" { return "ManagementAddress" }
    if yname == "chassis-id" { return "ChassisId" }
    return ""
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetSegmentPath() string {
    return "sender-id"
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "chassis-id" {
        return &senderId.ChassisId
    }
    return nil
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["chassis-id"] = &senderId.ChassisId
    return children
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["management-address-domain"] = senderId.ManagementAddressDomain
    leafs["management-address"] = senderId.ManagementAddress
    return leafs
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetBundleName() string { return "cisco_ios_xr" }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetYangName() string { return "sender-id" }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) SetParent(parent types.Entity) { senderId.parent = parent }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetParent() types.Entity { return senderId.parent }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetParentYangName() string { return "linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId
// Chassis ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Chassis ID Type. The type is CfmPmChassisIdFmt.
    ChassisIdType interface{}

    // Chassis ID Type. The type is interface{} with range: 0..255.
    ChassisIdTypeValue interface{}

    // Chassis ID (Deprecated). The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ChassisId interface{}

    // Chassis ID (Current).
    ChassisIdValue Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetFilter() yfilter.YFilter { return chassisId.YFilter }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) SetFilter(yf yfilter.YFilter) { chassisId.YFilter = yf }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetGoName(yname string) string {
    if yname == "chassis-id-type" { return "ChassisIdType" }
    if yname == "chassis-id-type-value" { return "ChassisIdTypeValue" }
    if yname == "chassis-id" { return "ChassisId" }
    if yname == "chassis-id-value" { return "ChassisIdValue" }
    return ""
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetSegmentPath() string {
    return "chassis-id"
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "chassis-id-value" {
        return &chassisId.ChassisIdValue
    }
    return nil
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["chassis-id-value"] = &chassisId.ChassisIdValue
    return children
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chassis-id-type"] = chassisId.ChassisIdType
    leafs["chassis-id-type-value"] = chassisId.ChassisIdTypeValue
    leafs["chassis-id"] = chassisId.ChassisId
    return leafs
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetBundleName() string { return "cisco_ios_xr" }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetYangName() string { return "chassis-id" }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) SetParent(parent types.Entity) { chassisId.parent = parent }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetParent() types.Entity { return chassisId.parent }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetParentYangName() string { return "sender-id" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue
// Chassis ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ChassisIDFormat. The type is CfmPmIdFmt.
    ChassisIdFormat interface{}

    // Chassis ID String. The type is string.
    ChassisIdString interface{}

    // Chassis ID MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    ChassisIdMac interface{}

    // Raw Chassis ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ChassisIdRaw interface{}
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetFilter() yfilter.YFilter { return chassisIdValue.YFilter }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) SetFilter(yf yfilter.YFilter) { chassisIdValue.YFilter = yf }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetGoName(yname string) string {
    if yname == "chassis-id-format" { return "ChassisIdFormat" }
    if yname == "chassis-id-string" { return "ChassisIdString" }
    if yname == "chassis-id-mac" { return "ChassisIdMac" }
    if yname == "chassis-id-raw" { return "ChassisIdRaw" }
    return ""
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetSegmentPath() string {
    return "chassis-id-value"
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chassis-id-format"] = chassisIdValue.ChassisIdFormat
    leafs["chassis-id-string"] = chassisIdValue.ChassisIdString
    leafs["chassis-id-mac"] = chassisIdValue.ChassisIdMac
    leafs["chassis-id-raw"] = chassisIdValue.ChassisIdRaw
    return leafs
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetBundleName() string { return "cisco_ios_xr" }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetYangName() string { return "chassis-id-value" }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) SetParent(parent types.Entity) { chassisIdValue.parent = parent }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetParent() types.Entity { return chassisIdValue.parent }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetParentYangName() string { return "chassis-id" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId
// Egress ID TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Last egress ID.
    LastEgressId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId

    // Next egress ID.
    NextEgressId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetFilter() yfilter.YFilter { return egressId.YFilter }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) SetFilter(yf yfilter.YFilter) { egressId.YFilter = yf }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetGoName(yname string) string {
    if yname == "last-egress-id" { return "LastEgressId" }
    if yname == "next-egress-id" { return "NextEgressId" }
    return ""
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetSegmentPath() string {
    return "egress-id"
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-egress-id" {
        return &egressId.LastEgressId
    }
    if childYangName == "next-egress-id" {
        return &egressId.NextEgressId
    }
    return nil
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-egress-id"] = &egressId.LastEgressId
    children["next-egress-id"] = &egressId.NextEgressId
    return children
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetBundleName() string { return "cisco_ios_xr" }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetYangName() string { return "egress-id" }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) SetParent(parent types.Entity) { egressId.parent = parent }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetParent() types.Entity { return egressId.parent }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetParentYangName() string { return "linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId
// Last egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetFilter() yfilter.YFilter { return lastEgressId.YFilter }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) SetFilter(yf yfilter.YFilter) { lastEgressId.YFilter = yf }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetGoName(yname string) string {
    if yname == "unique-id" { return "UniqueId" }
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetSegmentPath() string {
    return "last-egress-id"
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unique-id"] = lastEgressId.UniqueId
    leafs["mac-address"] = lastEgressId.MacAddress
    return leafs
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetBundleName() string { return "cisco_ios_xr" }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetYangName() string { return "last-egress-id" }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) SetParent(parent types.Entity) { lastEgressId.parent = parent }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetParent() types.Entity { return lastEgressId.parent }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetParentYangName() string { return "egress-id" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId
// Next egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetFilter() yfilter.YFilter { return nextEgressId.YFilter }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) SetFilter(yf yfilter.YFilter) { nextEgressId.YFilter = yf }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetGoName(yname string) string {
    if yname == "unique-id" { return "UniqueId" }
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetSegmentPath() string {
    return "next-egress-id"
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unique-id"] = nextEgressId.UniqueId
    leafs["mac-address"] = nextEgressId.MacAddress
    return leafs
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetBundleName() string { return "cisco_ios_xr" }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetYangName() string { return "next-egress-id" }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) SetParent(parent types.Entity) { nextEgressId.parent = parent }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetParent() types.Entity { return nextEgressId.parent }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetParentYangName() string { return "egress-id" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress
// Reply ingress TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reply ingress action. The type is CfmPmIngressAction.
    Action interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Port ID.
    PortId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetFilter() yfilter.YFilter { return replyIngress.YFilter }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) SetFilter(yf yfilter.YFilter) { replyIngress.YFilter = yf }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "port-id" { return "PortId" }
    return ""
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetSegmentPath() string {
    return "reply-ingress"
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-id" {
        return &replyIngress.PortId
    }
    return nil
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-id"] = &replyIngress.PortId
    return children
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = replyIngress.Action
    leafs["mac-address"] = replyIngress.MacAddress
    return leafs
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetBundleName() string { return "cisco_ios_xr" }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetYangName() string { return "reply-ingress" }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) SetParent(parent types.Entity) { replyIngress.parent = parent }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetParent() types.Entity { return replyIngress.parent }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetParentYangName() string { return "linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId
// Port ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port ID type. The type is CfmPmPortIdFmt.
    PortIdType interface{}

    // Port ID type value. The type is interface{} with range: 0..255.
    PortIdTypeValue interface{}

    // Port ID (Deprecated). The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    PortId interface{}

    // Port ID (Current).
    PortIdValue Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetFilter() yfilter.YFilter { return portId.YFilter }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) SetFilter(yf yfilter.YFilter) { portId.YFilter = yf }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetGoName(yname string) string {
    if yname == "port-id-type" { return "PortIdType" }
    if yname == "port-id-type-value" { return "PortIdTypeValue" }
    if yname == "port-id" { return "PortId" }
    if yname == "port-id-value" { return "PortIdValue" }
    return ""
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetSegmentPath() string {
    return "port-id"
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-id-value" {
        return &portId.PortIdValue
    }
    return nil
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-id-value"] = &portId.PortIdValue
    return children
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-id-type"] = portId.PortIdType
    leafs["port-id-type-value"] = portId.PortIdTypeValue
    leafs["port-id"] = portId.PortId
    return leafs
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetBundleName() string { return "cisco_ios_xr" }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetYangName() string { return "port-id" }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) SetParent(parent types.Entity) { portId.parent = parent }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetParent() types.Entity { return portId.parent }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetParentYangName() string { return "reply-ingress" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue
// Port ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PortIDFormat. The type is CfmPmIdFmt.
    PortIdFormat interface{}

    // Port ID String. The type is string.
    PortIdString interface{}

    // Port ID MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PortIdMac interface{}

    // Raw Port ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    PortIdRaw interface{}
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetFilter() yfilter.YFilter { return portIdValue.YFilter }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) SetFilter(yf yfilter.YFilter) { portIdValue.YFilter = yf }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetGoName(yname string) string {
    if yname == "port-id-format" { return "PortIdFormat" }
    if yname == "port-id-string" { return "PortIdString" }
    if yname == "port-id-mac" { return "PortIdMac" }
    if yname == "port-id-raw" { return "PortIdRaw" }
    return ""
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetSegmentPath() string {
    return "port-id-value"
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-id-format"] = portIdValue.PortIdFormat
    leafs["port-id-string"] = portIdValue.PortIdString
    leafs["port-id-mac"] = portIdValue.PortIdMac
    leafs["port-id-raw"] = portIdValue.PortIdRaw
    return leafs
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetBundleName() string { return "cisco_ios_xr" }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetYangName() string { return "port-id-value" }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) SetParent(parent types.Entity) { portIdValue.parent = parent }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetParent() types.Entity { return portIdValue.parent }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetParentYangName() string { return "port-id" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress
// Reply egress TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reply egress action. The type is CfmPmEgressAction.
    Action interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Port ID.
    PortId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetFilter() yfilter.YFilter { return replyEgress.YFilter }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) SetFilter(yf yfilter.YFilter) { replyEgress.YFilter = yf }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "port-id" { return "PortId" }
    return ""
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetSegmentPath() string {
    return "reply-egress"
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-id" {
        return &replyEgress.PortId
    }
    return nil
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-id"] = &replyEgress.PortId
    return children
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = replyEgress.Action
    leafs["mac-address"] = replyEgress.MacAddress
    return leafs
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetBundleName() string { return "cisco_ios_xr" }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetYangName() string { return "reply-egress" }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) SetParent(parent types.Entity) { replyEgress.parent = parent }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetParent() types.Entity { return replyEgress.parent }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetParentYangName() string { return "linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId
// Port ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port ID type. The type is CfmPmPortIdFmt.
    PortIdType interface{}

    // Port ID type value. The type is interface{} with range: 0..255.
    PortIdTypeValue interface{}

    // Port ID (Deprecated). The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    PortId interface{}

    // Port ID (Current).
    PortIdValue Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetFilter() yfilter.YFilter { return portId.YFilter }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) SetFilter(yf yfilter.YFilter) { portId.YFilter = yf }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetGoName(yname string) string {
    if yname == "port-id-type" { return "PortIdType" }
    if yname == "port-id-type-value" { return "PortIdTypeValue" }
    if yname == "port-id" { return "PortId" }
    if yname == "port-id-value" { return "PortIdValue" }
    return ""
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetSegmentPath() string {
    return "port-id"
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-id-value" {
        return &portId.PortIdValue
    }
    return nil
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-id-value"] = &portId.PortIdValue
    return children
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-id-type"] = portId.PortIdType
    leafs["port-id-type-value"] = portId.PortIdTypeValue
    leafs["port-id"] = portId.PortId
    return leafs
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetBundleName() string { return "cisco_ios_xr" }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetYangName() string { return "port-id" }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) SetParent(parent types.Entity) { portId.parent = parent }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetParent() types.Entity { return portId.parent }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetParentYangName() string { return "reply-egress" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue
// Port ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PortIDFormat. The type is CfmPmIdFmt.
    PortIdFormat interface{}

    // Port ID String. The type is string.
    PortIdString interface{}

    // Port ID MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PortIdMac interface{}

    // Raw Port ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    PortIdRaw interface{}
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetFilter() yfilter.YFilter { return portIdValue.YFilter }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) SetFilter(yf yfilter.YFilter) { portIdValue.YFilter = yf }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetGoName(yname string) string {
    if yname == "port-id-format" { return "PortIdFormat" }
    if yname == "port-id-string" { return "PortIdString" }
    if yname == "port-id-mac" { return "PortIdMac" }
    if yname == "port-id-raw" { return "PortIdRaw" }
    return ""
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetSegmentPath() string {
    return "port-id-value"
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-id-format"] = portIdValue.PortIdFormat
    leafs["port-id-string"] = portIdValue.PortIdString
    leafs["port-id-mac"] = portIdValue.PortIdMac
    leafs["port-id-raw"] = portIdValue.PortIdRaw
    return leafs
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetBundleName() string { return "cisco_ios_xr" }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetYangName() string { return "port-id-value" }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) SetParent(parent types.Entity) { portIdValue.parent = parent }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetParent() types.Entity { return portIdValue.parent }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetParentYangName() string { return "port-id" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop
// Last hop ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LastHopFormat. The type is CfmPmLastHopFmt.
    LastHopFormat interface{}

    // Hostname. The type is string.
    HostName interface{}

    // Egress ID.
    EgressId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetFilter() yfilter.YFilter { return lastHop.YFilter }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) SetFilter(yf yfilter.YFilter) { lastHop.YFilter = yf }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetGoName(yname string) string {
    if yname == "last-hop-format" { return "LastHopFormat" }
    if yname == "host-name" { return "HostName" }
    if yname == "egress-id" { return "EgressId" }
    return ""
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetSegmentPath() string {
    return "last-hop"
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "egress-id" {
        return &lastHop.EgressId
    }
    return nil
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["egress-id"] = &lastHop.EgressId
    return children
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["last-hop-format"] = lastHop.LastHopFormat
    leafs["host-name"] = lastHop.HostName
    return leafs
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetBundleName() string { return "cisco_ios_xr" }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetYangName() string { return "last-hop" }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) SetParent(parent types.Entity) { lastHop.parent = parent }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetParent() types.Entity { return lastHop.parent }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetParentYangName() string { return "linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId
// Egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetFilter() yfilter.YFilter { return egressId.YFilter }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) SetFilter(yf yfilter.YFilter) { egressId.YFilter = yf }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetGoName(yname string) string {
    if yname == "unique-id" { return "UniqueId" }
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetSegmentPath() string {
    return "egress-id"
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unique-id"] = egressId.UniqueId
    leafs["mac-address"] = egressId.MacAddress
    return leafs
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetBundleName() string { return "cisco_ios_xr" }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetYangName() string { return "egress-id" }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) SetParent(parent types.Entity) { egressId.parent = parent }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetParent() types.Entity { return egressId.parent }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetParentYangName() string { return "last-hop" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv
// Organizational-specific TLVs
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Organizationally-unique ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Oui interface{}

    // Subtype of TLV. The type is interface{} with range: 0..255.
    Subtype interface{}

    // Binary data in TLV. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetFilter() yfilter.YFilter { return organizationSpecificTlv.YFilter }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) SetFilter(yf yfilter.YFilter) { organizationSpecificTlv.YFilter = yf }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetGoName(yname string) string {
    if yname == "oui" { return "Oui" }
    if yname == "subtype" { return "Subtype" }
    if yname == "value" { return "Value" }
    return ""
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetSegmentPath() string {
    return "organization-specific-tlv"
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oui"] = organizationSpecificTlv.Oui
    leafs["subtype"] = organizationSpecificTlv.Subtype
    leafs["value"] = organizationSpecificTlv.Value
    return leafs
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetBundleName() string { return "cisco_ios_xr" }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetYangName() string { return "organization-specific-tlv" }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) SetParent(parent types.Entity) { organizationSpecificTlv.parent = parent }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetParent() types.Entity { return organizationSpecificTlv.parent }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetParentYangName() string { return "linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv
// Unknown TLVs
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type code of TLV. The type is interface{} with range: 0..255.
    Typecode interface{}

    // Binary data in TLV. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetFilter() yfilter.YFilter { return unknownTlv.YFilter }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) SetFilter(yf yfilter.YFilter) { unknownTlv.YFilter = yf }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetGoName(yname string) string {
    if yname == "typecode" { return "Typecode" }
    if yname == "value" { return "Value" }
    return ""
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetSegmentPath() string {
    return "unknown-tlv"
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["typecode"] = unknownTlv.Typecode
    leafs["value"] = unknownTlv.Value
    return leafs
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetBundleName() string { return "cisco_ios_xr" }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetYangName() string { return "unknown-tlv" }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) SetParent(parent types.Entity) { unknownTlv.parent = parent }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetParent() types.Entity { return unknownTlv.parent }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetParentYangName() string { return "linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply
// Received exploratory linktrace replies
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Undecoded frame. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    RawData interface{}

    // Frame header.
    Header Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header

    // Sender ID TLV.
    SenderId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId

    // Reply ingress TLV.
    ReplyIngress Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress

    // Reply egress TLV.
    ReplyEgress Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress

    // Last hop ID.
    LastHop Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop

    // Organizational-specific TLVs. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv.
    OrganizationSpecificTlv []Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv

    // Unknown TLVs. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv.
    UnknownTlv []Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv
}

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetFilter() yfilter.YFilter { return exploratoryLinktraceReply.YFilter }

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) SetFilter(yf yfilter.YFilter) { exploratoryLinktraceReply.YFilter = yf }

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetGoName(yname string) string {
    if yname == "raw-data" { return "RawData" }
    if yname == "header" { return "Header" }
    if yname == "sender-id" { return "SenderId" }
    if yname == "reply-ingress" { return "ReplyIngress" }
    if yname == "reply-egress" { return "ReplyEgress" }
    if yname == "last-hop" { return "LastHop" }
    if yname == "organization-specific-tlv" { return "OrganizationSpecificTlv" }
    if yname == "unknown-tlv" { return "UnknownTlv" }
    return ""
}

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetSegmentPath() string {
    return "exploratory-linktrace-reply"
}

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &exploratoryLinktraceReply.Header
    }
    if childYangName == "sender-id" {
        return &exploratoryLinktraceReply.SenderId
    }
    if childYangName == "reply-ingress" {
        return &exploratoryLinktraceReply.ReplyIngress
    }
    if childYangName == "reply-egress" {
        return &exploratoryLinktraceReply.ReplyEgress
    }
    if childYangName == "last-hop" {
        return &exploratoryLinktraceReply.LastHop
    }
    if childYangName == "organization-specific-tlv" {
        for _, c := range exploratoryLinktraceReply.OrganizationSpecificTlv {
            if exploratoryLinktraceReply.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv{}
        exploratoryLinktraceReply.OrganizationSpecificTlv = append(exploratoryLinktraceReply.OrganizationSpecificTlv, child)
        return &exploratoryLinktraceReply.OrganizationSpecificTlv[len(exploratoryLinktraceReply.OrganizationSpecificTlv)-1]
    }
    if childYangName == "unknown-tlv" {
        for _, c := range exploratoryLinktraceReply.UnknownTlv {
            if exploratoryLinktraceReply.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv{}
        exploratoryLinktraceReply.UnknownTlv = append(exploratoryLinktraceReply.UnknownTlv, child)
        return &exploratoryLinktraceReply.UnknownTlv[len(exploratoryLinktraceReply.UnknownTlv)-1]
    }
    return nil
}

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &exploratoryLinktraceReply.Header
    children["sender-id"] = &exploratoryLinktraceReply.SenderId
    children["reply-ingress"] = &exploratoryLinktraceReply.ReplyIngress
    children["reply-egress"] = &exploratoryLinktraceReply.ReplyEgress
    children["last-hop"] = &exploratoryLinktraceReply.LastHop
    for i := range exploratoryLinktraceReply.OrganizationSpecificTlv {
        children[exploratoryLinktraceReply.OrganizationSpecificTlv[i].GetSegmentPath()] = &exploratoryLinktraceReply.OrganizationSpecificTlv[i]
    }
    for i := range exploratoryLinktraceReply.UnknownTlv {
        children[exploratoryLinktraceReply.UnknownTlv[i].GetSegmentPath()] = &exploratoryLinktraceReply.UnknownTlv[i]
    }
    return children
}

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["raw-data"] = exploratoryLinktraceReply.RawData
    return leafs
}

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetBundleName() string { return "cisco_ios_xr" }

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetYangName() string { return "exploratory-linktrace-reply" }

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) SetParent(parent types.Entity) { exploratoryLinktraceReply.parent = parent }

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetParent() types.Entity { return exploratoryLinktraceReply.parent }

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetParentYangName() string { return "traceroute-cache" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header
// Frame header
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MD level. The type is CfmBagMdLevel.
    Level interface{}

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // ELR was forwarded. The type is bool.
    Forwarded interface{}

    // Terminal MEP reached. The type is bool.
    TerminalMep interface{}

    // Reply Filter unrecognized. The type is bool.
    ReplyFilterUnknown interface{}

    // Transaction ID. The type is interface{} with range: 0..4294967295.
    TransactionId interface{}

    // TTL. The type is interface{} with range: 0..255.
    Ttl interface{}

    // Relay action. The type is CfmPmElrRelayAction.
    RelayAction interface{}

    // Next Hop Timeout, in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    NextHopTimeout interface{}

    // Delay Model. The type is CfmPmEltDelayModel.
    DelayModel interface{}
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "version" { return "Version" }
    if yname == "forwarded" { return "Forwarded" }
    if yname == "terminal-mep" { return "TerminalMep" }
    if yname == "reply-filter-unknown" { return "ReplyFilterUnknown" }
    if yname == "transaction-id" { return "TransactionId" }
    if yname == "ttl" { return "Ttl" }
    if yname == "relay-action" { return "RelayAction" }
    if yname == "next-hop-timeout" { return "NextHopTimeout" }
    if yname == "delay-model" { return "DelayModel" }
    return ""
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetSegmentPath() string {
    return "header"
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = header.Level
    leafs["version"] = header.Version
    leafs["forwarded"] = header.Forwarded
    leafs["terminal-mep"] = header.TerminalMep
    leafs["reply-filter-unknown"] = header.ReplyFilterUnknown
    leafs["transaction-id"] = header.TransactionId
    leafs["ttl"] = header.Ttl
    leafs["relay-action"] = header.RelayAction
    leafs["next-hop-timeout"] = header.NextHopTimeout
    leafs["delay-model"] = header.DelayModel
    return leafs
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetYangName() string { return "header" }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetParent() types.Entity { return header.parent }

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetParentYangName() string { return "exploratory-linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId
// Sender ID TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Management address domain. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ManagementAddressDomain interface{}

    // Management address. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ManagementAddress interface{}

    // Chassis ID.
    ChassisId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetFilter() yfilter.YFilter { return senderId.YFilter }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) SetFilter(yf yfilter.YFilter) { senderId.YFilter = yf }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetGoName(yname string) string {
    if yname == "management-address-domain" { return "ManagementAddressDomain" }
    if yname == "management-address" { return "ManagementAddress" }
    if yname == "chassis-id" { return "ChassisId" }
    return ""
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetSegmentPath() string {
    return "sender-id"
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "chassis-id" {
        return &senderId.ChassisId
    }
    return nil
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["chassis-id"] = &senderId.ChassisId
    return children
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["management-address-domain"] = senderId.ManagementAddressDomain
    leafs["management-address"] = senderId.ManagementAddress
    return leafs
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetBundleName() string { return "cisco_ios_xr" }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetYangName() string { return "sender-id" }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) SetParent(parent types.Entity) { senderId.parent = parent }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetParent() types.Entity { return senderId.parent }

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetParentYangName() string { return "exploratory-linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId
// Chassis ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Chassis ID Type. The type is CfmPmChassisIdFmt.
    ChassisIdType interface{}

    // Chassis ID Type. The type is interface{} with range: 0..255.
    ChassisIdTypeValue interface{}

    // Chassis ID (Deprecated). The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ChassisId interface{}

    // Chassis ID (Current).
    ChassisIdValue Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetFilter() yfilter.YFilter { return chassisId.YFilter }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) SetFilter(yf yfilter.YFilter) { chassisId.YFilter = yf }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetGoName(yname string) string {
    if yname == "chassis-id-type" { return "ChassisIdType" }
    if yname == "chassis-id-type-value" { return "ChassisIdTypeValue" }
    if yname == "chassis-id" { return "ChassisId" }
    if yname == "chassis-id-value" { return "ChassisIdValue" }
    return ""
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetSegmentPath() string {
    return "chassis-id"
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "chassis-id-value" {
        return &chassisId.ChassisIdValue
    }
    return nil
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["chassis-id-value"] = &chassisId.ChassisIdValue
    return children
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chassis-id-type"] = chassisId.ChassisIdType
    leafs["chassis-id-type-value"] = chassisId.ChassisIdTypeValue
    leafs["chassis-id"] = chassisId.ChassisId
    return leafs
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetBundleName() string { return "cisco_ios_xr" }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetYangName() string { return "chassis-id" }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) SetParent(parent types.Entity) { chassisId.parent = parent }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetParent() types.Entity { return chassisId.parent }

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetParentYangName() string { return "sender-id" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue
// Chassis ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ChassisIDFormat. The type is CfmPmIdFmt.
    ChassisIdFormat interface{}

    // Chassis ID String. The type is string.
    ChassisIdString interface{}

    // Chassis ID MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    ChassisIdMac interface{}

    // Raw Chassis ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ChassisIdRaw interface{}
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetFilter() yfilter.YFilter { return chassisIdValue.YFilter }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) SetFilter(yf yfilter.YFilter) { chassisIdValue.YFilter = yf }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetGoName(yname string) string {
    if yname == "chassis-id-format" { return "ChassisIdFormat" }
    if yname == "chassis-id-string" { return "ChassisIdString" }
    if yname == "chassis-id-mac" { return "ChassisIdMac" }
    if yname == "chassis-id-raw" { return "ChassisIdRaw" }
    return ""
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetSegmentPath() string {
    return "chassis-id-value"
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chassis-id-format"] = chassisIdValue.ChassisIdFormat
    leafs["chassis-id-string"] = chassisIdValue.ChassisIdString
    leafs["chassis-id-mac"] = chassisIdValue.ChassisIdMac
    leafs["chassis-id-raw"] = chassisIdValue.ChassisIdRaw
    return leafs
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetBundleName() string { return "cisco_ios_xr" }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetYangName() string { return "chassis-id-value" }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) SetParent(parent types.Entity) { chassisIdValue.parent = parent }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetParent() types.Entity { return chassisIdValue.parent }

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetParentYangName() string { return "chassis-id" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress
// Reply ingress TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ELR Reply ingress action. The type is CfmPmElrIngressAction.
    Action interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Last egress ID.
    LastEgressId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId

    // Next egress ID.
    NextEgressId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId

    // Port ID.
    PortId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetFilter() yfilter.YFilter { return replyIngress.YFilter }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) SetFilter(yf yfilter.YFilter) { replyIngress.YFilter = yf }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "last-egress-id" { return "LastEgressId" }
    if yname == "next-egress-id" { return "NextEgressId" }
    if yname == "port-id" { return "PortId" }
    return ""
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetSegmentPath() string {
    return "reply-ingress"
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-egress-id" {
        return &replyIngress.LastEgressId
    }
    if childYangName == "next-egress-id" {
        return &replyIngress.NextEgressId
    }
    if childYangName == "port-id" {
        return &replyIngress.PortId
    }
    return nil
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-egress-id"] = &replyIngress.LastEgressId
    children["next-egress-id"] = &replyIngress.NextEgressId
    children["port-id"] = &replyIngress.PortId
    return children
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = replyIngress.Action
    leafs["mac-address"] = replyIngress.MacAddress
    return leafs
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetBundleName() string { return "cisco_ios_xr" }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetYangName() string { return "reply-ingress" }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) SetParent(parent types.Entity) { replyIngress.parent = parent }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetParent() types.Entity { return replyIngress.parent }

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetParentYangName() string { return "exploratory-linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId
// Last egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetFilter() yfilter.YFilter { return lastEgressId.YFilter }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) SetFilter(yf yfilter.YFilter) { lastEgressId.YFilter = yf }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetGoName(yname string) string {
    if yname == "unique-id" { return "UniqueId" }
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetSegmentPath() string {
    return "last-egress-id"
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unique-id"] = lastEgressId.UniqueId
    leafs["mac-address"] = lastEgressId.MacAddress
    return leafs
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetBundleName() string { return "cisco_ios_xr" }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetYangName() string { return "last-egress-id" }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) SetParent(parent types.Entity) { lastEgressId.parent = parent }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetParent() types.Entity { return lastEgressId.parent }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetParentYangName() string { return "reply-ingress" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId
// Next egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetFilter() yfilter.YFilter { return nextEgressId.YFilter }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) SetFilter(yf yfilter.YFilter) { nextEgressId.YFilter = yf }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetGoName(yname string) string {
    if yname == "unique-id" { return "UniqueId" }
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetSegmentPath() string {
    return "next-egress-id"
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unique-id"] = nextEgressId.UniqueId
    leafs["mac-address"] = nextEgressId.MacAddress
    return leafs
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetBundleName() string { return "cisco_ios_xr" }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetYangName() string { return "next-egress-id" }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) SetParent(parent types.Entity) { nextEgressId.parent = parent }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetParent() types.Entity { return nextEgressId.parent }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetParentYangName() string { return "reply-ingress" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId
// Port ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port ID type. The type is CfmPmPortIdFmt.
    PortIdType interface{}

    // Port ID type value. The type is interface{} with range: 0..255.
    PortIdTypeValue interface{}

    // Port ID (Deprecated). The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    PortId interface{}

    // Port ID (Current).
    PortIdValue Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetFilter() yfilter.YFilter { return portId.YFilter }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) SetFilter(yf yfilter.YFilter) { portId.YFilter = yf }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetGoName(yname string) string {
    if yname == "port-id-type" { return "PortIdType" }
    if yname == "port-id-type-value" { return "PortIdTypeValue" }
    if yname == "port-id" { return "PortId" }
    if yname == "port-id-value" { return "PortIdValue" }
    return ""
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetSegmentPath() string {
    return "port-id"
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-id-value" {
        return &portId.PortIdValue
    }
    return nil
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-id-value"] = &portId.PortIdValue
    return children
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-id-type"] = portId.PortIdType
    leafs["port-id-type-value"] = portId.PortIdTypeValue
    leafs["port-id"] = portId.PortId
    return leafs
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetBundleName() string { return "cisco_ios_xr" }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetYangName() string { return "port-id" }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) SetParent(parent types.Entity) { portId.parent = parent }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetParent() types.Entity { return portId.parent }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetParentYangName() string { return "reply-ingress" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue
// Port ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PortIDFormat. The type is CfmPmIdFmt.
    PortIdFormat interface{}

    // Port ID String. The type is string.
    PortIdString interface{}

    // Port ID MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PortIdMac interface{}

    // Raw Port ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    PortIdRaw interface{}
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetFilter() yfilter.YFilter { return portIdValue.YFilter }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) SetFilter(yf yfilter.YFilter) { portIdValue.YFilter = yf }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetGoName(yname string) string {
    if yname == "port-id-format" { return "PortIdFormat" }
    if yname == "port-id-string" { return "PortIdString" }
    if yname == "port-id-mac" { return "PortIdMac" }
    if yname == "port-id-raw" { return "PortIdRaw" }
    return ""
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetSegmentPath() string {
    return "port-id-value"
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-id-format"] = portIdValue.PortIdFormat
    leafs["port-id-string"] = portIdValue.PortIdString
    leafs["port-id-mac"] = portIdValue.PortIdMac
    leafs["port-id-raw"] = portIdValue.PortIdRaw
    return leafs
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetBundleName() string { return "cisco_ios_xr" }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetYangName() string { return "port-id-value" }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) SetParent(parent types.Entity) { portIdValue.parent = parent }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetParent() types.Entity { return portIdValue.parent }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetParentYangName() string { return "port-id" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress
// Reply egress TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reply egress action. The type is CfmPmElrEgressAction.
    Action interface{}

    // MAC address of egress interface. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Last Egress ID.
    LastEgressId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId

    // Next Egress ID.
    NextEgressId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId

    // Port ID.
    PortId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetFilter() yfilter.YFilter { return replyEgress.YFilter }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) SetFilter(yf yfilter.YFilter) { replyEgress.YFilter = yf }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "last-egress-id" { return "LastEgressId" }
    if yname == "next-egress-id" { return "NextEgressId" }
    if yname == "port-id" { return "PortId" }
    return ""
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetSegmentPath() string {
    return "reply-egress"
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-egress-id" {
        return &replyEgress.LastEgressId
    }
    if childYangName == "next-egress-id" {
        return &replyEgress.NextEgressId
    }
    if childYangName == "port-id" {
        return &replyEgress.PortId
    }
    return nil
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-egress-id"] = &replyEgress.LastEgressId
    children["next-egress-id"] = &replyEgress.NextEgressId
    children["port-id"] = &replyEgress.PortId
    return children
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = replyEgress.Action
    leafs["mac-address"] = replyEgress.MacAddress
    return leafs
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetBundleName() string { return "cisco_ios_xr" }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetYangName() string { return "reply-egress" }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) SetParent(parent types.Entity) { replyEgress.parent = parent }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetParent() types.Entity { return replyEgress.parent }

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetParentYangName() string { return "exploratory-linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId
// Last Egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetFilter() yfilter.YFilter { return lastEgressId.YFilter }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) SetFilter(yf yfilter.YFilter) { lastEgressId.YFilter = yf }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetGoName(yname string) string {
    if yname == "unique-id" { return "UniqueId" }
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetSegmentPath() string {
    return "last-egress-id"
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unique-id"] = lastEgressId.UniqueId
    leafs["mac-address"] = lastEgressId.MacAddress
    return leafs
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetBundleName() string { return "cisco_ios_xr" }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetYangName() string { return "last-egress-id" }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) SetParent(parent types.Entity) { lastEgressId.parent = parent }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetParent() types.Entity { return lastEgressId.parent }

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetParentYangName() string { return "reply-egress" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId
// Next Egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetFilter() yfilter.YFilter { return nextEgressId.YFilter }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) SetFilter(yf yfilter.YFilter) { nextEgressId.YFilter = yf }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetGoName(yname string) string {
    if yname == "unique-id" { return "UniqueId" }
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetSegmentPath() string {
    return "next-egress-id"
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unique-id"] = nextEgressId.UniqueId
    leafs["mac-address"] = nextEgressId.MacAddress
    return leafs
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetBundleName() string { return "cisco_ios_xr" }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetYangName() string { return "next-egress-id" }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) SetParent(parent types.Entity) { nextEgressId.parent = parent }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetParent() types.Entity { return nextEgressId.parent }

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetParentYangName() string { return "reply-egress" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId
// Port ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port ID type. The type is CfmPmPortIdFmt.
    PortIdType interface{}

    // Port ID type value. The type is interface{} with range: 0..255.
    PortIdTypeValue interface{}

    // Port ID (Deprecated). The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    PortId interface{}

    // Port ID (Current).
    PortIdValue Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetFilter() yfilter.YFilter { return portId.YFilter }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) SetFilter(yf yfilter.YFilter) { portId.YFilter = yf }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetGoName(yname string) string {
    if yname == "port-id-type" { return "PortIdType" }
    if yname == "port-id-type-value" { return "PortIdTypeValue" }
    if yname == "port-id" { return "PortId" }
    if yname == "port-id-value" { return "PortIdValue" }
    return ""
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetSegmentPath() string {
    return "port-id"
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-id-value" {
        return &portId.PortIdValue
    }
    return nil
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-id-value"] = &portId.PortIdValue
    return children
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-id-type"] = portId.PortIdType
    leafs["port-id-type-value"] = portId.PortIdTypeValue
    leafs["port-id"] = portId.PortId
    return leafs
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetBundleName() string { return "cisco_ios_xr" }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetYangName() string { return "port-id" }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) SetParent(parent types.Entity) { portId.parent = parent }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetParent() types.Entity { return portId.parent }

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetParentYangName() string { return "reply-egress" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue
// Port ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PortIDFormat. The type is CfmPmIdFmt.
    PortIdFormat interface{}

    // Port ID String. The type is string.
    PortIdString interface{}

    // Port ID MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PortIdMac interface{}

    // Raw Port ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    PortIdRaw interface{}
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetFilter() yfilter.YFilter { return portIdValue.YFilter }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) SetFilter(yf yfilter.YFilter) { portIdValue.YFilter = yf }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetGoName(yname string) string {
    if yname == "port-id-format" { return "PortIdFormat" }
    if yname == "port-id-string" { return "PortIdString" }
    if yname == "port-id-mac" { return "PortIdMac" }
    if yname == "port-id-raw" { return "PortIdRaw" }
    return ""
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetSegmentPath() string {
    return "port-id-value"
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-id-format"] = portIdValue.PortIdFormat
    leafs["port-id-string"] = portIdValue.PortIdString
    leafs["port-id-mac"] = portIdValue.PortIdMac
    leafs["port-id-raw"] = portIdValue.PortIdRaw
    return leafs
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetBundleName() string { return "cisco_ios_xr" }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetYangName() string { return "port-id-value" }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) SetParent(parent types.Entity) { portIdValue.parent = parent }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetParent() types.Entity { return portIdValue.parent }

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetParentYangName() string { return "port-id" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop
// Last hop ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LastHopFormat. The type is CfmPmLastHopFmt.
    LastHopFormat interface{}

    // Hostname. The type is string.
    HostName interface{}

    // Egress ID.
    EgressId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetFilter() yfilter.YFilter { return lastHop.YFilter }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) SetFilter(yf yfilter.YFilter) { lastHop.YFilter = yf }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetGoName(yname string) string {
    if yname == "last-hop-format" { return "LastHopFormat" }
    if yname == "host-name" { return "HostName" }
    if yname == "egress-id" { return "EgressId" }
    return ""
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetSegmentPath() string {
    return "last-hop"
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "egress-id" {
        return &lastHop.EgressId
    }
    return nil
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["egress-id"] = &lastHop.EgressId
    return children
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["last-hop-format"] = lastHop.LastHopFormat
    leafs["host-name"] = lastHop.HostName
    return leafs
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetBundleName() string { return "cisco_ios_xr" }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetYangName() string { return "last-hop" }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) SetParent(parent types.Entity) { lastHop.parent = parent }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetParent() types.Entity { return lastHop.parent }

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetParentYangName() string { return "exploratory-linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId
// Egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetFilter() yfilter.YFilter { return egressId.YFilter }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) SetFilter(yf yfilter.YFilter) { egressId.YFilter = yf }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetGoName(yname string) string {
    if yname == "unique-id" { return "UniqueId" }
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetSegmentPath() string {
    return "egress-id"
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unique-id"] = egressId.UniqueId
    leafs["mac-address"] = egressId.MacAddress
    return leafs
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetBundleName() string { return "cisco_ios_xr" }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetYangName() string { return "egress-id" }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) SetParent(parent types.Entity) { egressId.parent = parent }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetParent() types.Entity { return egressId.parent }

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetParentYangName() string { return "last-hop" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv
// Organizational-specific TLVs
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Organizationally-unique ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Oui interface{}

    // Subtype of TLV. The type is interface{} with range: 0..255.
    Subtype interface{}

    // Binary data in TLV. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetFilter() yfilter.YFilter { return organizationSpecificTlv.YFilter }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) SetFilter(yf yfilter.YFilter) { organizationSpecificTlv.YFilter = yf }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetGoName(yname string) string {
    if yname == "oui" { return "Oui" }
    if yname == "subtype" { return "Subtype" }
    if yname == "value" { return "Value" }
    return ""
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetSegmentPath() string {
    return "organization-specific-tlv"
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oui"] = organizationSpecificTlv.Oui
    leafs["subtype"] = organizationSpecificTlv.Subtype
    leafs["value"] = organizationSpecificTlv.Value
    return leafs
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetBundleName() string { return "cisco_ios_xr" }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetYangName() string { return "organization-specific-tlv" }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) SetParent(parent types.Entity) { organizationSpecificTlv.parent = parent }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetParent() types.Entity { return organizationSpecificTlv.parent }

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetParentYangName() string { return "exploratory-linktrace-reply" }

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv
// Unknown TLVs
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type code of TLV. The type is interface{} with range: 0..255.
    Typecode interface{}

    // Binary data in TLV. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetFilter() yfilter.YFilter { return unknownTlv.YFilter }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) SetFilter(yf yfilter.YFilter) { unknownTlv.YFilter = yf }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetGoName(yname string) string {
    if yname == "typecode" { return "Typecode" }
    if yname == "value" { return "Value" }
    return ""
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetSegmentPath() string {
    return "unknown-tlv"
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["typecode"] = unknownTlv.Typecode
    leafs["value"] = unknownTlv.Value
    return leafs
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetBundleName() string { return "cisco_ios_xr" }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetYangName() string { return "unknown-tlv" }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) SetParent(parent types.Entity) { unknownTlv.parent = parent }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetParent() types.Entity { return unknownTlv.parent }

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetParentYangName() string { return "exploratory-linktrace-reply" }

// Cfm_Global_LocalMeps
// Local MEPs table
type Cfm_Global_LocalMeps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular local MEP. The type is slice of
    // Cfm_Global_LocalMeps_LocalMep.
    LocalMep []Cfm_Global_LocalMeps_LocalMep
}

func (localMeps *Cfm_Global_LocalMeps) GetFilter() yfilter.YFilter { return localMeps.YFilter }

func (localMeps *Cfm_Global_LocalMeps) SetFilter(yf yfilter.YFilter) { localMeps.YFilter = yf }

func (localMeps *Cfm_Global_LocalMeps) GetGoName(yname string) string {
    if yname == "local-mep" { return "LocalMep" }
    return ""
}

func (localMeps *Cfm_Global_LocalMeps) GetSegmentPath() string {
    return "local-meps"
}

func (localMeps *Cfm_Global_LocalMeps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-mep" {
        for _, c := range localMeps.LocalMep {
            if localMeps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_LocalMeps_LocalMep{}
        localMeps.LocalMep = append(localMeps.LocalMep, child)
        return &localMeps.LocalMep[len(localMeps.LocalMep)-1]
    }
    return nil
}

func (localMeps *Cfm_Global_LocalMeps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localMeps.LocalMep {
        children[localMeps.LocalMep[i].GetSegmentPath()] = &localMeps.LocalMep[i]
    }
    return children
}

func (localMeps *Cfm_Global_LocalMeps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localMeps *Cfm_Global_LocalMeps) GetBundleName() string { return "cisco_ios_xr" }

func (localMeps *Cfm_Global_LocalMeps) GetYangName() string { return "local-meps" }

func (localMeps *Cfm_Global_LocalMeps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localMeps *Cfm_Global_LocalMeps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localMeps *Cfm_Global_LocalMeps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localMeps *Cfm_Global_LocalMeps) SetParent(parent types.Entity) { localMeps.parent = parent }

func (localMeps *Cfm_Global_LocalMeps) GetParent() types.Entity { return localMeps.parent }

func (localMeps *Cfm_Global_LocalMeps) GetParentYangName() string { return "global" }

// Cfm_Global_LocalMeps_LocalMep
// Information about a particular local MEP
type Cfm_Global_LocalMeps_LocalMep struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. MEP ID. The type is interface{} with range:
    // 1..8191.
    MepId interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}

    // Maintenance domain name. The type is string.
    DomainXr interface{}

    // Service name. The type is string.
    ServiceXr interface{}

    // Maintenance level. The type is CfmBagMdLevel.
    Level interface{}

    // MEP ID. The type is interface{} with range: 0..65535.
    MepIdXr interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceXr interface{}

    // IM Interface state. The type is string.
    InterfaceState interface{}

    // Interface interworking state. The type is CfmBagIwState.
    InterworkingState interface{}

    // STP state. The type is CfmBagStpState.
    StpState interface{}

    // MEP facing direction. The type is CfmBagDirection.
    MepDirection interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Number of peer MEPs detected. The type is interface{} with range:
    // 0..4294967295.
    PeerMepsDetected interface{}

    // Number of peer MEPs detected with errors. The type is interface{} with
    // range: 0..4294967295.
    PeerMepsWithErrorsDetected interface{}

    // Remote defect indicated. The type is bool.
    RemoteDefect interface{}

    // Fault Notification Generation state. The type is CfmPmMepFngState.
    FaultNotificationState interface{}

    // CCM generation enabled. The type is bool.
    CcmGenerationEnabled interface{}

    // The interval between CCMs. The type is CfmBagCcmInterval.
    CcmInterval interface{}

    // Offload status of CCM processing. The type is CfmBagCcmOffload.
    CcmOffload interface{}

    // Highest-priority defect present since last FNG reset. The type is
    // CfmPmMepDefect.
    HighestDefect interface{}

    // A peer MEP RDI defect has been reported. The type is bool.
    RdiDefect interface{}

    // A peer MEP port or interface status error has been reported. The type is
    // bool.
    MacStatusDefect interface{}

    // A peer MEP CCM error has been reported. The type is bool.
    PeerMepCcmDefect interface{}

    // A CCM error has been reported. The type is bool.
    ErrorCcmDefect interface{}

    // A cross-connect CCM error has been reported. The type is bool.
    CrossConnectCcmDefect interface{}

    // Next Transaction ID to be sent in a Loopback Message. The type is
    // interface{} with range: 0..4294967295.
    NextLbmId interface{}

    // Next Transaction ID to be sent in a Linktrace Message. The type is
    // interface{} with range: 0..4294967295.
    NextLtmId interface{}

    // CoS bits the MEP will use for sent packets, if configured. The type is
    // interface{} with range: 0..255.
    Cos interface{}

    // EFD triggered on the interface. The type is bool.
    EfdTriggered interface{}

    // The local MEP is on an interface in standby mode. The type is bool.
    Standby interface{}

    // MEP is on a sub-interface in the same bridge-domain and on the same trunk
    // interface as another sub-interface. The type is bool.
    Hairpin interface{}

    // Defects present but ignored due to 'report defects' configuration. The type
    // is bool.
    DefectsIgnored interface{}

    // MEP statistics.
    Statistics Cfm_Global_LocalMeps_LocalMep_Statistics

    // MEP AIS statistics.
    AisStatistics Cfm_Global_LocalMeps_LocalMep_AisStatistics

    // Defects detected from peer MEPs.
    Defects Cfm_Global_LocalMeps_LocalMep_Defects
}

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetFilter() yfilter.YFilter { return localMep.YFilter }

func (localMep *Cfm_Global_LocalMeps_LocalMep) SetFilter(yf yfilter.YFilter) { localMep.YFilter = yf }

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "service" { return "Service" }
    if yname == "mep-id" { return "MepId" }
    if yname == "interface" { return "Interface" }
    if yname == "domain-xr" { return "DomainXr" }
    if yname == "service-xr" { return "ServiceXr" }
    if yname == "level" { return "Level" }
    if yname == "mep-id-xr" { return "MepIdXr" }
    if yname == "interface-xr" { return "InterfaceXr" }
    if yname == "interface-state" { return "InterfaceState" }
    if yname == "interworking-state" { return "InterworkingState" }
    if yname == "stp-state" { return "StpState" }
    if yname == "mep-direction" { return "MepDirection" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "peer-meps-detected" { return "PeerMepsDetected" }
    if yname == "peer-meps-with-errors-detected" { return "PeerMepsWithErrorsDetected" }
    if yname == "remote-defect" { return "RemoteDefect" }
    if yname == "fault-notification-state" { return "FaultNotificationState" }
    if yname == "ccm-generation-enabled" { return "CcmGenerationEnabled" }
    if yname == "ccm-interval" { return "CcmInterval" }
    if yname == "ccm-offload" { return "CcmOffload" }
    if yname == "highest-defect" { return "HighestDefect" }
    if yname == "rdi-defect" { return "RdiDefect" }
    if yname == "mac-status-defect" { return "MacStatusDefect" }
    if yname == "peer-mep-ccm-defect" { return "PeerMepCcmDefect" }
    if yname == "error-ccm-defect" { return "ErrorCcmDefect" }
    if yname == "cross-connect-ccm-defect" { return "CrossConnectCcmDefect" }
    if yname == "next-lbm-id" { return "NextLbmId" }
    if yname == "next-ltm-id" { return "NextLtmId" }
    if yname == "cos" { return "Cos" }
    if yname == "efd-triggered" { return "EfdTriggered" }
    if yname == "standby" { return "Standby" }
    if yname == "hairpin" { return "Hairpin" }
    if yname == "defects-ignored" { return "DefectsIgnored" }
    if yname == "statistics" { return "Statistics" }
    if yname == "ais-statistics" { return "AisStatistics" }
    if yname == "defects" { return "Defects" }
    return ""
}

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetSegmentPath() string {
    return "local-mep" + "[domain='" + fmt.Sprintf("%v", localMep.Domain) + "']" + "[service='" + fmt.Sprintf("%v", localMep.Service) + "']" + "[mep-id='" + fmt.Sprintf("%v", localMep.MepId) + "']" + "[interface='" + fmt.Sprintf("%v", localMep.Interface) + "']"
}

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &localMep.Statistics
    }
    if childYangName == "ais-statistics" {
        return &localMep.AisStatistics
    }
    if childYangName == "defects" {
        return &localMep.Defects
    }
    return nil
}

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &localMep.Statistics
    children["ais-statistics"] = &localMep.AisStatistics
    children["defects"] = &localMep.Defects
    return children
}

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = localMep.Domain
    leafs["service"] = localMep.Service
    leafs["mep-id"] = localMep.MepId
    leafs["interface"] = localMep.Interface
    leafs["domain-xr"] = localMep.DomainXr
    leafs["service-xr"] = localMep.ServiceXr
    leafs["level"] = localMep.Level
    leafs["mep-id-xr"] = localMep.MepIdXr
    leafs["interface-xr"] = localMep.InterfaceXr
    leafs["interface-state"] = localMep.InterfaceState
    leafs["interworking-state"] = localMep.InterworkingState
    leafs["stp-state"] = localMep.StpState
    leafs["mep-direction"] = localMep.MepDirection
    leafs["mac-address"] = localMep.MacAddress
    leafs["peer-meps-detected"] = localMep.PeerMepsDetected
    leafs["peer-meps-with-errors-detected"] = localMep.PeerMepsWithErrorsDetected
    leafs["remote-defect"] = localMep.RemoteDefect
    leafs["fault-notification-state"] = localMep.FaultNotificationState
    leafs["ccm-generation-enabled"] = localMep.CcmGenerationEnabled
    leafs["ccm-interval"] = localMep.CcmInterval
    leafs["ccm-offload"] = localMep.CcmOffload
    leafs["highest-defect"] = localMep.HighestDefect
    leafs["rdi-defect"] = localMep.RdiDefect
    leafs["mac-status-defect"] = localMep.MacStatusDefect
    leafs["peer-mep-ccm-defect"] = localMep.PeerMepCcmDefect
    leafs["error-ccm-defect"] = localMep.ErrorCcmDefect
    leafs["cross-connect-ccm-defect"] = localMep.CrossConnectCcmDefect
    leafs["next-lbm-id"] = localMep.NextLbmId
    leafs["next-ltm-id"] = localMep.NextLtmId
    leafs["cos"] = localMep.Cos
    leafs["efd-triggered"] = localMep.EfdTriggered
    leafs["standby"] = localMep.Standby
    leafs["hairpin"] = localMep.Hairpin
    leafs["defects-ignored"] = localMep.DefectsIgnored
    return leafs
}

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetBundleName() string { return "cisco_ios_xr" }

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetYangName() string { return "local-mep" }

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localMep *Cfm_Global_LocalMeps_LocalMep) SetParent(parent types.Entity) { localMep.parent = parent }

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetParent() types.Entity { return localMep.parent }

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetParentYangName() string { return "local-meps" }

// Cfm_Global_LocalMeps_LocalMep_Statistics
// MEP statistics
type Cfm_Global_LocalMeps_LocalMep_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of CCMs sent. The type is interface{} with range:
    // 0..18446744073709551615.
    CcmsSent interface{}

    // Number of CCMs received. The type is interface{} with range:
    // 0..18446744073709551615.
    CcmsReceived interface{}

    // Number of CCMs received out-of-sequence. The type is interface{} with
    // range: 0..18446744073709551615.
    CcmsOutOfSequence interface{}

    // Number of CCMs discarded because maximum MEPs limit was reached. The type
    // is interface{} with range: 0..18446744073709551615.
    CcmsDiscarded interface{}

    // Number of LBMs sent. The type is interface{} with range:
    // 0..18446744073709551615.
    LbMsSent interface{}

    // Number of LBRs sent. The type is interface{} with range:
    // 0..18446744073709551615.
    LbRsSent interface{}

    // Number of LBRs received. The type is interface{} with range:
    // 0..18446744073709551615.
    LbRsReceived interface{}

    // Number of LBRs received out-of-sequence. The type is interface{} with
    // range: 0..18446744073709551615.
    LbRsOutOfSequence interface{}

    // Number of LBRs received with non-matching user-specified data. The type is
    // interface{} with range: 0..18446744073709551615.
    LbRsBadData interface{}

    // Number of LBMs received. The type is interface{} with range:
    // 0..18446744073709551615.
    LbMsReceived interface{}

    // Number of unexpected LTRs received. The type is interface{} with range:
    // 0..18446744073709551615.
    LtRsReceivedUnexpected interface{}

    // Number of AIS messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AiSsSent interface{}

    // Number of AIS messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    AiSsReceived interface{}

    // Number of LCK messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    LcKsReceived interface{}

    // Number of DMM messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    DmMsSent interface{}

    // Number of DMM messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    DmMsReceived interface{}

    // Number of DMR messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    DmRsSent interface{}

    // Number of DMR messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    DmRsReceived interface{}

    // Number of SLM messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    SlMsSent interface{}

    // Number of SLM messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    SlMsReceived interface{}

    // Number of SLR messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    SlRsSent interface{}

    // Number of SLR messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    SlRsReceived interface{}

    // Number of LMM messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    LmMsSent interface{}

    // Number of LMM messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    LmMsReceived interface{}

    // Number of LMR messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    LmRsSent interface{}

    // Number of LMR messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    LmRsReceived interface{}

    // Number of BNM messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    BnMsReceived interface{}

    // Number of BNM messages discarded. The type is interface{} with range:
    // 0..18446744073709551615.
    BnMsDiscarded interface{}
}

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetGoName(yname string) string {
    if yname == "ccms-sent" { return "CcmsSent" }
    if yname == "ccms-received" { return "CcmsReceived" }
    if yname == "ccms-out-of-sequence" { return "CcmsOutOfSequence" }
    if yname == "ccms-discarded" { return "CcmsDiscarded" }
    if yname == "lb-ms-sent" { return "LbMsSent" }
    if yname == "lb-rs-sent" { return "LbRsSent" }
    if yname == "lb-rs-received" { return "LbRsReceived" }
    if yname == "lb-rs-out-of-sequence" { return "LbRsOutOfSequence" }
    if yname == "lb-rs-bad-data" { return "LbRsBadData" }
    if yname == "lb-ms-received" { return "LbMsReceived" }
    if yname == "lt-rs-received-unexpected" { return "LtRsReceivedUnexpected" }
    if yname == "ai-ss-sent" { return "AiSsSent" }
    if yname == "ai-ss-received" { return "AiSsReceived" }
    if yname == "lc-ks-received" { return "LcKsReceived" }
    if yname == "dm-ms-sent" { return "DmMsSent" }
    if yname == "dm-ms-received" { return "DmMsReceived" }
    if yname == "dm-rs-sent" { return "DmRsSent" }
    if yname == "dm-rs-received" { return "DmRsReceived" }
    if yname == "sl-ms-sent" { return "SlMsSent" }
    if yname == "sl-ms-received" { return "SlMsReceived" }
    if yname == "sl-rs-sent" { return "SlRsSent" }
    if yname == "sl-rs-received" { return "SlRsReceived" }
    if yname == "lm-ms-sent" { return "LmMsSent" }
    if yname == "lm-ms-received" { return "LmMsReceived" }
    if yname == "lm-rs-sent" { return "LmRsSent" }
    if yname == "lm-rs-received" { return "LmRsReceived" }
    if yname == "bn-ms-received" { return "BnMsReceived" }
    if yname == "bn-ms-discarded" { return "BnMsDiscarded" }
    return ""
}

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccms-sent"] = statistics.CcmsSent
    leafs["ccms-received"] = statistics.CcmsReceived
    leafs["ccms-out-of-sequence"] = statistics.CcmsOutOfSequence
    leafs["ccms-discarded"] = statistics.CcmsDiscarded
    leafs["lb-ms-sent"] = statistics.LbMsSent
    leafs["lb-rs-sent"] = statistics.LbRsSent
    leafs["lb-rs-received"] = statistics.LbRsReceived
    leafs["lb-rs-out-of-sequence"] = statistics.LbRsOutOfSequence
    leafs["lb-rs-bad-data"] = statistics.LbRsBadData
    leafs["lb-ms-received"] = statistics.LbMsReceived
    leafs["lt-rs-received-unexpected"] = statistics.LtRsReceivedUnexpected
    leafs["ai-ss-sent"] = statistics.AiSsSent
    leafs["ai-ss-received"] = statistics.AiSsReceived
    leafs["lc-ks-received"] = statistics.LcKsReceived
    leafs["dm-ms-sent"] = statistics.DmMsSent
    leafs["dm-ms-received"] = statistics.DmMsReceived
    leafs["dm-rs-sent"] = statistics.DmRsSent
    leafs["dm-rs-received"] = statistics.DmRsReceived
    leafs["sl-ms-sent"] = statistics.SlMsSent
    leafs["sl-ms-received"] = statistics.SlMsReceived
    leafs["sl-rs-sent"] = statistics.SlRsSent
    leafs["sl-rs-received"] = statistics.SlRsReceived
    leafs["lm-ms-sent"] = statistics.LmMsSent
    leafs["lm-ms-received"] = statistics.LmMsReceived
    leafs["lm-rs-sent"] = statistics.LmRsSent
    leafs["lm-rs-received"] = statistics.LmRsReceived
    leafs["bn-ms-received"] = statistics.BnMsReceived
    leafs["bn-ms-discarded"] = statistics.BnMsDiscarded
    return leafs
}

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetYangName() string { return "statistics" }

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetParentYangName() string { return "local-mep" }

// Cfm_Global_LocalMeps_LocalMep_AisStatistics
// MEP AIS statistics
type Cfm_Global_LocalMeps_LocalMep_AisStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AIS transmission level. The type is CfmBagMdLevel.
    Level interface{}

    // AIS transmission interval. The type is CfmBagAisInterval.
    Interval interface{}

    // Details of how AIS is being transmitted. The type is CfmPmAisTransmit.
    SendingAis interface{}

    // Details of how the signal is being received. The type is CfmPmAisReceive.
    ReceivingAis interface{}

    // The interval of the last received AIS packet. The type is
    // CfmBagAisInterval.
    LastInterval interface{}

    // Source MAC address of the last received AIS packet. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    LastMacAddress interface{}

    // Time elapsed since AIS sending started.
    SendingStart Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart

    // Time elapsed since AIS receiving started.
    ReceivingStart Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart
}

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetFilter() yfilter.YFilter { return aisStatistics.YFilter }

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) SetFilter(yf yfilter.YFilter) { aisStatistics.YFilter = yf }

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "interval" { return "Interval" }
    if yname == "sending-ais" { return "SendingAis" }
    if yname == "receiving-ais" { return "ReceivingAis" }
    if yname == "last-interval" { return "LastInterval" }
    if yname == "last-mac-address" { return "LastMacAddress" }
    if yname == "sending-start" { return "SendingStart" }
    if yname == "receiving-start" { return "ReceivingStart" }
    return ""
}

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetSegmentPath() string {
    return "ais-statistics"
}

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sending-start" {
        return &aisStatistics.SendingStart
    }
    if childYangName == "receiving-start" {
        return &aisStatistics.ReceivingStart
    }
    return nil
}

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sending-start"] = &aisStatistics.SendingStart
    children["receiving-start"] = &aisStatistics.ReceivingStart
    return children
}

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = aisStatistics.Level
    leafs["interval"] = aisStatistics.Interval
    leafs["sending-ais"] = aisStatistics.SendingAis
    leafs["receiving-ais"] = aisStatistics.ReceivingAis
    leafs["last-interval"] = aisStatistics.LastInterval
    leafs["last-mac-address"] = aisStatistics.LastMacAddress
    return leafs
}

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetYangName() string { return "ais-statistics" }

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) SetParent(parent types.Entity) { aisStatistics.parent = parent }

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetParent() types.Entity { return aisStatistics.parent }

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetParentYangName() string { return "local-mep" }

// Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart
// Time elapsed since AIS sending started
type Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetFilter() yfilter.YFilter { return sendingStart.YFilter }

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) SetFilter(yf yfilter.YFilter) { sendingStart.YFilter = yf }

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetSegmentPath() string {
    return "sending-start"
}

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = sendingStart.Seconds
    leafs["nanoseconds"] = sendingStart.Nanoseconds
    return leafs
}

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetBundleName() string { return "cisco_ios_xr" }

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetYangName() string { return "sending-start" }

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) SetParent(parent types.Entity) { sendingStart.parent = parent }

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetParent() types.Entity { return sendingStart.parent }

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetParentYangName() string { return "ais-statistics" }

// Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart
// Time elapsed since AIS receiving started
type Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetFilter() yfilter.YFilter { return receivingStart.YFilter }

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) SetFilter(yf yfilter.YFilter) { receivingStart.YFilter = yf }

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetSegmentPath() string {
    return "receiving-start"
}

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = receivingStart.Seconds
    leafs["nanoseconds"] = receivingStart.Nanoseconds
    return leafs
}

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetBundleName() string { return "cisco_ios_xr" }

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetYangName() string { return "receiving-start" }

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) SetParent(parent types.Entity) { receivingStart.parent = parent }

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetParent() types.Entity { return receivingStart.parent }

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetParentYangName() string { return "ais-statistics" }

// Cfm_Global_LocalMeps_LocalMep_Defects
// Defects detected from peer MEPs
type Cfm_Global_LocalMeps_LocalMep_Defects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AIS or LCK received. The type is bool.
    AisReceived interface{}

    // Number of peer MEPs that have timed out. The type is interface{} with
    // range: 0..4294967295.
    PeerMepsThatTimedOut interface{}

    // Number of missing peer MEPs. The type is interface{} with range:
    // 0..4294967295.
    Missing interface{}

    // Number of missing auto cross-check MEPs. The type is interface{} with
    // range: 0..4294967295.
    AutoMissing interface{}

    // Number of unexpected peer MEPs. The type is interface{} with range:
    // 0..4294967295.
    Unexpected interface{}

    // The local port or interface is down. The type is bool.
    LocalPortStatus interface{}

    // A peer port or interface is down. The type is bool.
    PeerPortStatus interface{}

    // Defects detected from remote MEPs.
    RemoteMepsDefects Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects
}

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetFilter() yfilter.YFilter { return defects.YFilter }

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) SetFilter(yf yfilter.YFilter) { defects.YFilter = yf }

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetGoName(yname string) string {
    if yname == "ais-received" { return "AisReceived" }
    if yname == "peer-meps-that-timed-out" { return "PeerMepsThatTimedOut" }
    if yname == "missing" { return "Missing" }
    if yname == "auto-missing" { return "AutoMissing" }
    if yname == "unexpected" { return "Unexpected" }
    if yname == "local-port-status" { return "LocalPortStatus" }
    if yname == "peer-port-status" { return "PeerPortStatus" }
    if yname == "remote-meps-defects" { return "RemoteMepsDefects" }
    return ""
}

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetSegmentPath() string {
    return "defects"
}

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-meps-defects" {
        return &defects.RemoteMepsDefects
    }
    return nil
}

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote-meps-defects"] = &defects.RemoteMepsDefects
    return children
}

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ais-received"] = defects.AisReceived
    leafs["peer-meps-that-timed-out"] = defects.PeerMepsThatTimedOut
    leafs["missing"] = defects.Missing
    leafs["auto-missing"] = defects.AutoMissing
    leafs["unexpected"] = defects.Unexpected
    leafs["local-port-status"] = defects.LocalPortStatus
    leafs["peer-port-status"] = defects.PeerPortStatus
    return leafs
}

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetBundleName() string { return "cisco_ios_xr" }

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetYangName() string { return "defects" }

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) SetParent(parent types.Entity) { defects.parent = parent }

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetParent() types.Entity { return defects.parent }

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetParentYangName() string { return "local-mep" }

// Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects
// Defects detected from remote MEPs
type Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timed out (loss threshold exceeded). The type is bool.
    LossThresholdExceeded interface{}

    // Invalid level. The type is bool.
    InvalidLevel interface{}

    // Invalid MAID. The type is bool.
    InvalidMaid interface{}

    // Invalid CCM interval. The type is bool.
    InvalidCcmInterval interface{}

    // Loop detected (our MAC address received). The type is bool.
    ReceivedOurMac interface{}

    // Configuration Error (our MEP ID received). The type is bool.
    ReceivedOurMepId interface{}

    // Remote defection indication received. The type is bool.
    ReceivedRdi interface{}
}

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetFilter() yfilter.YFilter { return remoteMepsDefects.YFilter }

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) SetFilter(yf yfilter.YFilter) { remoteMepsDefects.YFilter = yf }

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetGoName(yname string) string {
    if yname == "loss-threshold-exceeded" { return "LossThresholdExceeded" }
    if yname == "invalid-level" { return "InvalidLevel" }
    if yname == "invalid-maid" { return "InvalidMaid" }
    if yname == "invalid-ccm-interval" { return "InvalidCcmInterval" }
    if yname == "received-our-mac" { return "ReceivedOurMac" }
    if yname == "received-our-mep-id" { return "ReceivedOurMepId" }
    if yname == "received-rdi" { return "ReceivedRdi" }
    return ""
}

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetSegmentPath() string {
    return "remote-meps-defects"
}

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["loss-threshold-exceeded"] = remoteMepsDefects.LossThresholdExceeded
    leafs["invalid-level"] = remoteMepsDefects.InvalidLevel
    leafs["invalid-maid"] = remoteMepsDefects.InvalidMaid
    leafs["invalid-ccm-interval"] = remoteMepsDefects.InvalidCcmInterval
    leafs["received-our-mac"] = remoteMepsDefects.ReceivedOurMac
    leafs["received-our-mep-id"] = remoteMepsDefects.ReceivedOurMepId
    leafs["received-rdi"] = remoteMepsDefects.ReceivedRdi
    return leafs
}

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetBundleName() string { return "cisco_ios_xr" }

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetYangName() string { return "remote-meps-defects" }

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) SetParent(parent types.Entity) { remoteMepsDefects.parent = parent }

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetParent() types.Entity { return remoteMepsDefects.parent }

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetParentYangName() string { return "defects" }

// Cfm_Global_PeerMePv2S
// Peer MEPs table Version 2
type Cfm_Global_PeerMePv2S struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a peer MEP for a particular local MEP. The type is slice
    // of Cfm_Global_PeerMePv2S_PeerMePv2.
    PeerMePv2 []Cfm_Global_PeerMePv2S_PeerMePv2
}

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetFilter() yfilter.YFilter { return peerMePv2S.YFilter }

func (peerMePv2S *Cfm_Global_PeerMePv2S) SetFilter(yf yfilter.YFilter) { peerMePv2S.YFilter = yf }

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetGoName(yname string) string {
    if yname == "peer-me-pv2" { return "PeerMePv2" }
    return ""
}

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetSegmentPath() string {
    return "peer-me-pv2s"
}

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-me-pv2" {
        for _, c := range peerMePv2S.PeerMePv2 {
            if peerMePv2S.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_PeerMePv2S_PeerMePv2{}
        peerMePv2S.PeerMePv2 = append(peerMePv2S.PeerMePv2, child)
        return &peerMePv2S.PeerMePv2[len(peerMePv2S.PeerMePv2)-1]
    }
    return nil
}

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerMePv2S.PeerMePv2 {
        children[peerMePv2S.PeerMePv2[i].GetSegmentPath()] = &peerMePv2S.PeerMePv2[i]
    }
    return children
}

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetBundleName() string { return "cisco_ios_xr" }

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetYangName() string { return "peer-me-pv2s" }

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerMePv2S *Cfm_Global_PeerMePv2S) SetParent(parent types.Entity) { peerMePv2S.parent = parent }

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetParent() types.Entity { return peerMePv2S.parent }

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetParentYangName() string { return "global" }

// Cfm_Global_PeerMePv2S_PeerMePv2
// Information about a peer MEP for a particular
// local MEP
type Cfm_Global_PeerMePv2S_PeerMePv2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. MEP ID of Local MEP. The type is interface{} with
    // range: 1..8191.
    LocalMepId interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}

    // This attribute is a key. MEP ID of Peer MEP. The type is interface{} with
    // range: 1..8191.
    PeerMepId interface{}

    // This attribute is a key. Peer MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PeerMacAddress interface{}

    // Maintenance domain name. The type is string.
    DomainXr interface{}

    // Service name. The type is string.
    ServiceXr interface{}

    // Maintenance level. The type is CfmBagMdLevel.
    Level interface{}

    // MEP ID. The type is interface{} with range: 0..65535.
    MepId interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceXr interface{}

    // MEP facing direction. The type is CfmBagDirection.
    MepDirection interface{}

    // The local MEP is on an interface in standby mode. The type is bool.
    Standby interface{}

    // Peer MEP.
    PeerMep Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep
}

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetFilter() yfilter.YFilter { return peerMePv2.YFilter }

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) SetFilter(yf yfilter.YFilter) { peerMePv2.YFilter = yf }

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "service" { return "Service" }
    if yname == "local-mep-id" { return "LocalMepId" }
    if yname == "interface" { return "Interface" }
    if yname == "peer-mep-id" { return "PeerMepId" }
    if yname == "peer-mac-address" { return "PeerMacAddress" }
    if yname == "domain-xr" { return "DomainXr" }
    if yname == "service-xr" { return "ServiceXr" }
    if yname == "level" { return "Level" }
    if yname == "mep-id" { return "MepId" }
    if yname == "interface-xr" { return "InterfaceXr" }
    if yname == "mep-direction" { return "MepDirection" }
    if yname == "standby" { return "Standby" }
    if yname == "peer-mep" { return "PeerMep" }
    return ""
}

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetSegmentPath() string {
    return "peer-me-pv2" + "[domain='" + fmt.Sprintf("%v", peerMePv2.Domain) + "']" + "[service='" + fmt.Sprintf("%v", peerMePv2.Service) + "']" + "[local-mep-id='" + fmt.Sprintf("%v", peerMePv2.LocalMepId) + "']" + "[interface='" + fmt.Sprintf("%v", peerMePv2.Interface) + "']" + "[peer-mep-id='" + fmt.Sprintf("%v", peerMePv2.PeerMepId) + "']" + "[peer-mac-address='" + fmt.Sprintf("%v", peerMePv2.PeerMacAddress) + "']"
}

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-mep" {
        return &peerMePv2.PeerMep
    }
    return nil
}

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-mep"] = &peerMePv2.PeerMep
    return children
}

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = peerMePv2.Domain
    leafs["service"] = peerMePv2.Service
    leafs["local-mep-id"] = peerMePv2.LocalMepId
    leafs["interface"] = peerMePv2.Interface
    leafs["peer-mep-id"] = peerMePv2.PeerMepId
    leafs["peer-mac-address"] = peerMePv2.PeerMacAddress
    leafs["domain-xr"] = peerMePv2.DomainXr
    leafs["service-xr"] = peerMePv2.ServiceXr
    leafs["level"] = peerMePv2.Level
    leafs["mep-id"] = peerMePv2.MepId
    leafs["interface-xr"] = peerMePv2.InterfaceXr
    leafs["mep-direction"] = peerMePv2.MepDirection
    leafs["standby"] = peerMePv2.Standby
    return leafs
}

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetBundleName() string { return "cisco_ios_xr" }

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetYangName() string { return "peer-me-pv2" }

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) SetParent(parent types.Entity) { peerMePv2.parent = parent }

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetParent() types.Entity { return peerMePv2.parent }

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetParentYangName() string { return "peer-me-pv2s" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep
// Peer MEP
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MEP ID. The type is interface{} with range: 0..65535.
    MepId interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Cross-check state. The type is CfmPmRmepXcState.
    CrossCheckState interface{}

    // State of the peer MEP state machine. The type is CfmPmRmepState.
    PeerMepState interface{}

    // Offload status of received CCM handling. The type is CfmBagCcmOffload.
    CcmOffload interface{}

    // Error state.
    ErrorState Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState

    // Elapsed time since peer MEP became active or timed out.
    LastUpDownTime Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime

    // Last CCM received from the peer MEP.
    LastCcmReceived Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived

    // Peer MEP statistics.
    Statistics Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics
}

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetFilter() yfilter.YFilter { return peerMep.YFilter }

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) SetFilter(yf yfilter.YFilter) { peerMep.YFilter = yf }

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetGoName(yname string) string {
    if yname == "mep-id" { return "MepId" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "cross-check-state" { return "CrossCheckState" }
    if yname == "peer-mep-state" { return "PeerMepState" }
    if yname == "ccm-offload" { return "CcmOffload" }
    if yname == "error-state" { return "ErrorState" }
    if yname == "last-up-down-time" { return "LastUpDownTime" }
    if yname == "last-ccm-received" { return "LastCcmReceived" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetSegmentPath() string {
    return "peer-mep"
}

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error-state" {
        return &peerMep.ErrorState
    }
    if childYangName == "last-up-down-time" {
        return &peerMep.LastUpDownTime
    }
    if childYangName == "last-ccm-received" {
        return &peerMep.LastCcmReceived
    }
    if childYangName == "statistics" {
        return &peerMep.Statistics
    }
    return nil
}

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["error-state"] = &peerMep.ErrorState
    children["last-up-down-time"] = &peerMep.LastUpDownTime
    children["last-ccm-received"] = &peerMep.LastCcmReceived
    children["statistics"] = &peerMep.Statistics
    return children
}

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mep-id"] = peerMep.MepId
    leafs["mac-address"] = peerMep.MacAddress
    leafs["cross-check-state"] = peerMep.CrossCheckState
    leafs["peer-mep-state"] = peerMep.PeerMepState
    leafs["ccm-offload"] = peerMep.CcmOffload
    return leafs
}

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetBundleName() string { return "cisco_ios_xr" }

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetYangName() string { return "peer-mep" }

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) SetParent(parent types.Entity) { peerMep.parent = parent }

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetParent() types.Entity { return peerMep.parent }

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetParentYangName() string { return "peer-me-pv2" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState
// Error state
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timed out (loss threshold exceeded). The type is bool.
    LossThresholdExceeded interface{}

    // Invalid level. The type is bool.
    InvalidLevel interface{}

    // Invalid MAID. The type is bool.
    InvalidMaid interface{}

    // Invalid CCM interval. The type is bool.
    InvalidCcmInterval interface{}

    // Loop detected (our MAC address received). The type is bool.
    ReceivedOurMac interface{}

    // Configuration Error (our MEP ID received). The type is bool.
    ReceivedOurMepId interface{}

    // Remote defection indication received. The type is bool.
    ReceivedRdi interface{}
}

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetFilter() yfilter.YFilter { return errorState.YFilter }

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) SetFilter(yf yfilter.YFilter) { errorState.YFilter = yf }

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetGoName(yname string) string {
    if yname == "loss-threshold-exceeded" { return "LossThresholdExceeded" }
    if yname == "invalid-level" { return "InvalidLevel" }
    if yname == "invalid-maid" { return "InvalidMaid" }
    if yname == "invalid-ccm-interval" { return "InvalidCcmInterval" }
    if yname == "received-our-mac" { return "ReceivedOurMac" }
    if yname == "received-our-mep-id" { return "ReceivedOurMepId" }
    if yname == "received-rdi" { return "ReceivedRdi" }
    return ""
}

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetSegmentPath() string {
    return "error-state"
}

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["loss-threshold-exceeded"] = errorState.LossThresholdExceeded
    leafs["invalid-level"] = errorState.InvalidLevel
    leafs["invalid-maid"] = errorState.InvalidMaid
    leafs["invalid-ccm-interval"] = errorState.InvalidCcmInterval
    leafs["received-our-mac"] = errorState.ReceivedOurMac
    leafs["received-our-mep-id"] = errorState.ReceivedOurMepId
    leafs["received-rdi"] = errorState.ReceivedRdi
    return leafs
}

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetBundleName() string { return "cisco_ios_xr" }

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetYangName() string { return "error-state" }

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) SetParent(parent types.Entity) { errorState.parent = parent }

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetParent() types.Entity { return errorState.parent }

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetParentYangName() string { return "peer-mep" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime
// Elapsed time since peer MEP became active or
// timed out
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetFilter() yfilter.YFilter { return lastUpDownTime.YFilter }

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) SetFilter(yf yfilter.YFilter) { lastUpDownTime.YFilter = yf }

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetSegmentPath() string {
    return "last-up-down-time"
}

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = lastUpDownTime.Seconds
    leafs["nanoseconds"] = lastUpDownTime.Nanoseconds
    return leafs
}

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetBundleName() string { return "cisco_ios_xr" }

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetYangName() string { return "last-up-down-time" }

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) SetParent(parent types.Entity) { lastUpDownTime.parent = parent }

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetParent() types.Entity { return lastUpDownTime.parent }

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetParentYangName() string { return "peer-mep" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived
// Last CCM received from the peer MEP
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port status. The type is CfmPmPortStatus.
    PortStatus interface{}

    // Interface status. The type is CfmPmIntfStatus.
    InterfaceStatus interface{}

    // Additional interface status. The type is CfmPmAddlIntfStatus.
    AdditionalInterfaceStatus interface{}

    // Undecoded frame. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    RawData interface{}

    // Frame header.
    Header Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header

    // Sender ID TLV.
    SenderId Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId

    // MEP name.
    MepName Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName

    // Organizational-specific TLVs. The type is slice of
    // Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv.
    OrganizationSpecificTlv []Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv

    // Unknown TLVs. The type is slice of
    // Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv.
    UnknownTlv []Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv
}

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetFilter() yfilter.YFilter { return lastCcmReceived.YFilter }

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) SetFilter(yf yfilter.YFilter) { lastCcmReceived.YFilter = yf }

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetGoName(yname string) string {
    if yname == "port-status" { return "PortStatus" }
    if yname == "interface-status" { return "InterfaceStatus" }
    if yname == "additional-interface-status" { return "AdditionalInterfaceStatus" }
    if yname == "raw-data" { return "RawData" }
    if yname == "header" { return "Header" }
    if yname == "sender-id" { return "SenderId" }
    if yname == "mep-name" { return "MepName" }
    if yname == "organization-specific-tlv" { return "OrganizationSpecificTlv" }
    if yname == "unknown-tlv" { return "UnknownTlv" }
    return ""
}

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetSegmentPath() string {
    return "last-ccm-received"
}

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &lastCcmReceived.Header
    }
    if childYangName == "sender-id" {
        return &lastCcmReceived.SenderId
    }
    if childYangName == "mep-name" {
        return &lastCcmReceived.MepName
    }
    if childYangName == "organization-specific-tlv" {
        for _, c := range lastCcmReceived.OrganizationSpecificTlv {
            if lastCcmReceived.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv{}
        lastCcmReceived.OrganizationSpecificTlv = append(lastCcmReceived.OrganizationSpecificTlv, child)
        return &lastCcmReceived.OrganizationSpecificTlv[len(lastCcmReceived.OrganizationSpecificTlv)-1]
    }
    if childYangName == "unknown-tlv" {
        for _, c := range lastCcmReceived.UnknownTlv {
            if lastCcmReceived.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv{}
        lastCcmReceived.UnknownTlv = append(lastCcmReceived.UnknownTlv, child)
        return &lastCcmReceived.UnknownTlv[len(lastCcmReceived.UnknownTlv)-1]
    }
    return nil
}

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &lastCcmReceived.Header
    children["sender-id"] = &lastCcmReceived.SenderId
    children["mep-name"] = &lastCcmReceived.MepName
    for i := range lastCcmReceived.OrganizationSpecificTlv {
        children[lastCcmReceived.OrganizationSpecificTlv[i].GetSegmentPath()] = &lastCcmReceived.OrganizationSpecificTlv[i]
    }
    for i := range lastCcmReceived.UnknownTlv {
        children[lastCcmReceived.UnknownTlv[i].GetSegmentPath()] = &lastCcmReceived.UnknownTlv[i]
    }
    return children
}

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-status"] = lastCcmReceived.PortStatus
    leafs["interface-status"] = lastCcmReceived.InterfaceStatus
    leafs["additional-interface-status"] = lastCcmReceived.AdditionalInterfaceStatus
    leafs["raw-data"] = lastCcmReceived.RawData
    return leafs
}

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetBundleName() string { return "cisco_ios_xr" }

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetYangName() string { return "last-ccm-received" }

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) SetParent(parent types.Entity) { lastCcmReceived.parent = parent }

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetParent() types.Entity { return lastCcmReceived.parent }

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetParentYangName() string { return "peer-mep" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header
// Frame header
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MD level. The type is CfmBagMdLevel.
    Level interface{}

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // CCM interval. The type is CfmBagCcmInterval.
    Interval interface{}

    // Remote defect indicated. The type is bool.
    Rdi interface{}

    // CCM sequence number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // MEP ID. The type is interface{} with range: 0..65535.
    MepId interface{}

    // MDID Format. The type is interface{} with range: 0..255.
    MdidFormat interface{}

    // Short MA Name format. The type is interface{} with range: 0..255.
    ShortMaNameFormat interface{}

    // MDID.
    Mdid Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid

    // Short MA Name.
    ShortMaName Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName
}

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "version" { return "Version" }
    if yname == "interval" { return "Interval" }
    if yname == "rdi" { return "Rdi" }
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "mep-id" { return "MepId" }
    if yname == "mdid-format" { return "MdidFormat" }
    if yname == "short-ma-name-format" { return "ShortMaNameFormat" }
    if yname == "mdid" { return "Mdid" }
    if yname == "short-ma-name" { return "ShortMaName" }
    return ""
}

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetSegmentPath() string {
    return "header"
}

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mdid" {
        return &header.Mdid
    }
    if childYangName == "short-ma-name" {
        return &header.ShortMaName
    }
    return nil
}

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mdid"] = &header.Mdid
    children["short-ma-name"] = &header.ShortMaName
    return children
}

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = header.Level
    leafs["version"] = header.Version
    leafs["interval"] = header.Interval
    leafs["rdi"] = header.Rdi
    leafs["sequence-number"] = header.SequenceNumber
    leafs["mep-id"] = header.MepId
    leafs["mdid-format"] = header.MdidFormat
    leafs["short-ma-name-format"] = header.ShortMaNameFormat
    return leafs
}

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetYangName() string { return "header" }

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetParent() types.Entity { return header.parent }

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetParentYangName() string { return "last-ccm-received" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid
// MDID
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MDIDFormatValue. The type is CfmBagMdidFmt.
    MdidFormatValue interface{}

    // DNS-like name. The type is string.
    DnsLikeName interface{}

    // String name. The type is string.
    StringName interface{}

    // Hex data. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    MdidData interface{}

    // MAC address name.
    MacName Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName
}

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetFilter() yfilter.YFilter { return mdid.YFilter }

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) SetFilter(yf yfilter.YFilter) { mdid.YFilter = yf }

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetGoName(yname string) string {
    if yname == "mdid-format-value" { return "MdidFormatValue" }
    if yname == "dns-like-name" { return "DnsLikeName" }
    if yname == "string-name" { return "StringName" }
    if yname == "mdid-data" { return "MdidData" }
    if yname == "mac-name" { return "MacName" }
    return ""
}

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetSegmentPath() string {
    return "mdid"
}

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-name" {
        return &mdid.MacName
    }
    return nil
}

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-name"] = &mdid.MacName
    return children
}

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mdid-format-value"] = mdid.MdidFormatValue
    leafs["dns-like-name"] = mdid.DnsLikeName
    leafs["string-name"] = mdid.StringName
    leafs["mdid-data"] = mdid.MdidData
    return leafs
}

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetBundleName() string { return "cisco_ios_xr" }

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetYangName() string { return "mdid" }

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) SetParent(parent types.Entity) { mdid.parent = parent }

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetParent() types.Entity { return mdid.parent }

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetParentYangName() string { return "header" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName
// MAC address name
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Integer. The type is interface{} with range: 0..65535.
    Integer interface{}
}

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetFilter() yfilter.YFilter { return macName.YFilter }

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) SetFilter(yf yfilter.YFilter) { macName.YFilter = yf }

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "integer" { return "Integer" }
    return ""
}

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetSegmentPath() string {
    return "mac-name"
}

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = macName.MacAddress
    leafs["integer"] = macName.Integer
    return leafs
}

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetBundleName() string { return "cisco_ios_xr" }

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetYangName() string { return "mac-name" }

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) SetParent(parent types.Entity) { macName.parent = parent }

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetParent() types.Entity { return macName.parent }

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetParentYangName() string { return "mdid" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName
// Short MA Name
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ShortMANameFormatValue. The type is CfmBagSmanFmt.
    ShortMaNameFormatValue interface{}

    // VLAN ID name. The type is interface{} with range: 0..65535.
    VlanIdName interface{}

    // String name. The type is string.
    StringName interface{}

    // Unsigned integer name. The type is interface{} with range: 0..65535.
    IntegerName interface{}

    // ICC-based format. The type is string.
    IccBased interface{}

    // Hex data. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ShortMaNameData interface{}

    // VPN ID name.
    VpnIdName Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName
}

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetFilter() yfilter.YFilter { return shortMaName.YFilter }

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) SetFilter(yf yfilter.YFilter) { shortMaName.YFilter = yf }

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetGoName(yname string) string {
    if yname == "short-ma-name-format-value" { return "ShortMaNameFormatValue" }
    if yname == "vlan-id-name" { return "VlanIdName" }
    if yname == "string-name" { return "StringName" }
    if yname == "integer-name" { return "IntegerName" }
    if yname == "icc-based" { return "IccBased" }
    if yname == "short-ma-name-data" { return "ShortMaNameData" }
    if yname == "vpn-id-name" { return "VpnIdName" }
    return ""
}

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetSegmentPath() string {
    return "short-ma-name"
}

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vpn-id-name" {
        return &shortMaName.VpnIdName
    }
    return nil
}

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vpn-id-name"] = &shortMaName.VpnIdName
    return children
}

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["short-ma-name-format-value"] = shortMaName.ShortMaNameFormatValue
    leafs["vlan-id-name"] = shortMaName.VlanIdName
    leafs["string-name"] = shortMaName.StringName
    leafs["integer-name"] = shortMaName.IntegerName
    leafs["icc-based"] = shortMaName.IccBased
    leafs["short-ma-name-data"] = shortMaName.ShortMaNameData
    return leafs
}

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetBundleName() string { return "cisco_ios_xr" }

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetYangName() string { return "short-ma-name" }

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) SetParent(parent types.Entity) { shortMaName.parent = parent }

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetParent() types.Entity { return shortMaName.parent }

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetParentYangName() string { return "header" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName
// VPN ID name
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VPN authority organizationally-unique ID. The type is interface{} with
    // range: 0..4294967295.
    Oui interface{}

    // VPN index. The type is interface{} with range: 0..4294967295.
    Index interface{}
}

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetFilter() yfilter.YFilter { return vpnIdName.YFilter }

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) SetFilter(yf yfilter.YFilter) { vpnIdName.YFilter = yf }

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetGoName(yname string) string {
    if yname == "oui" { return "Oui" }
    if yname == "index" { return "Index" }
    return ""
}

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetSegmentPath() string {
    return "vpn-id-name"
}

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oui"] = vpnIdName.Oui
    leafs["index"] = vpnIdName.Index
    return leafs
}

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetBundleName() string { return "cisco_ios_xr" }

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetYangName() string { return "vpn-id-name" }

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) SetParent(parent types.Entity) { vpnIdName.parent = parent }

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetParent() types.Entity { return vpnIdName.parent }

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetParentYangName() string { return "short-ma-name" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId
// Sender ID TLV
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Management address domain. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ManagementAddressDomain interface{}

    // Management address. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ManagementAddress interface{}

    // Chassis ID.
    ChassisId Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId
}

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetFilter() yfilter.YFilter { return senderId.YFilter }

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) SetFilter(yf yfilter.YFilter) { senderId.YFilter = yf }

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetGoName(yname string) string {
    if yname == "management-address-domain" { return "ManagementAddressDomain" }
    if yname == "management-address" { return "ManagementAddress" }
    if yname == "chassis-id" { return "ChassisId" }
    return ""
}

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetSegmentPath() string {
    return "sender-id"
}

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "chassis-id" {
        return &senderId.ChassisId
    }
    return nil
}

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["chassis-id"] = &senderId.ChassisId
    return children
}

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["management-address-domain"] = senderId.ManagementAddressDomain
    leafs["management-address"] = senderId.ManagementAddress
    return leafs
}

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetBundleName() string { return "cisco_ios_xr" }

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetYangName() string { return "sender-id" }

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) SetParent(parent types.Entity) { senderId.parent = parent }

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetParent() types.Entity { return senderId.parent }

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetParentYangName() string { return "last-ccm-received" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId
// Chassis ID
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Chassis ID Type. The type is CfmPmChassisIdFmt.
    ChassisIdType interface{}

    // Chassis ID Type. The type is interface{} with range: 0..255.
    ChassisIdTypeValue interface{}

    // Chassis ID (Deprecated). The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ChassisId interface{}

    // Chassis ID (Current).
    ChassisIdValue Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue
}

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetFilter() yfilter.YFilter { return chassisId.YFilter }

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) SetFilter(yf yfilter.YFilter) { chassisId.YFilter = yf }

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetGoName(yname string) string {
    if yname == "chassis-id-type" { return "ChassisIdType" }
    if yname == "chassis-id-type-value" { return "ChassisIdTypeValue" }
    if yname == "chassis-id" { return "ChassisId" }
    if yname == "chassis-id-value" { return "ChassisIdValue" }
    return ""
}

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetSegmentPath() string {
    return "chassis-id"
}

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "chassis-id-value" {
        return &chassisId.ChassisIdValue
    }
    return nil
}

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["chassis-id-value"] = &chassisId.ChassisIdValue
    return children
}

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chassis-id-type"] = chassisId.ChassisIdType
    leafs["chassis-id-type-value"] = chassisId.ChassisIdTypeValue
    leafs["chassis-id"] = chassisId.ChassisId
    return leafs
}

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetBundleName() string { return "cisco_ios_xr" }

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetYangName() string { return "chassis-id" }

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) SetParent(parent types.Entity) { chassisId.parent = parent }

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetParent() types.Entity { return chassisId.parent }

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetParentYangName() string { return "sender-id" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue
// Chassis ID (Current)
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ChassisIDFormat. The type is CfmPmIdFmt.
    ChassisIdFormat interface{}

    // Chassis ID String. The type is string.
    ChassisIdString interface{}

    // Chassis ID MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    ChassisIdMac interface{}

    // Raw Chassis ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    ChassisIdRaw interface{}
}

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetFilter() yfilter.YFilter { return chassisIdValue.YFilter }

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) SetFilter(yf yfilter.YFilter) { chassisIdValue.YFilter = yf }

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetGoName(yname string) string {
    if yname == "chassis-id-format" { return "ChassisIdFormat" }
    if yname == "chassis-id-string" { return "ChassisIdString" }
    if yname == "chassis-id-mac" { return "ChassisIdMac" }
    if yname == "chassis-id-raw" { return "ChassisIdRaw" }
    return ""
}

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetSegmentPath() string {
    return "chassis-id-value"
}

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chassis-id-format"] = chassisIdValue.ChassisIdFormat
    leafs["chassis-id-string"] = chassisIdValue.ChassisIdString
    leafs["chassis-id-mac"] = chassisIdValue.ChassisIdMac
    leafs["chassis-id-raw"] = chassisIdValue.ChassisIdRaw
    return leafs
}

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetBundleName() string { return "cisco_ios_xr" }

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetYangName() string { return "chassis-id-value" }

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) SetParent(parent types.Entity) { chassisIdValue.parent = parent }

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetParent() types.Entity { return chassisIdValue.parent }

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetParentYangName() string { return "chassis-id" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName
// MEP name
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MEP name. The type is string.
    Name interface{}
}

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetFilter() yfilter.YFilter { return mepName.YFilter }

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) SetFilter(yf yfilter.YFilter) { mepName.YFilter = yf }

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetSegmentPath() string {
    return "mep-name"
}

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = mepName.Name
    return leafs
}

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetBundleName() string { return "cisco_ios_xr" }

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetYangName() string { return "mep-name" }

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) SetParent(parent types.Entity) { mepName.parent = parent }

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetParent() types.Entity { return mepName.parent }

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetParentYangName() string { return "last-ccm-received" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv
// Organizational-specific TLVs
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Organizationally-unique ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Oui interface{}

    // Subtype of TLV. The type is interface{} with range: 0..255.
    Subtype interface{}

    // Binary data in TLV. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetFilter() yfilter.YFilter { return organizationSpecificTlv.YFilter }

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) SetFilter(yf yfilter.YFilter) { organizationSpecificTlv.YFilter = yf }

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetGoName(yname string) string {
    if yname == "oui" { return "Oui" }
    if yname == "subtype" { return "Subtype" }
    if yname == "value" { return "Value" }
    return ""
}

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetSegmentPath() string {
    return "organization-specific-tlv"
}

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oui"] = organizationSpecificTlv.Oui
    leafs["subtype"] = organizationSpecificTlv.Subtype
    leafs["value"] = organizationSpecificTlv.Value
    return leafs
}

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetBundleName() string { return "cisco_ios_xr" }

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetYangName() string { return "organization-specific-tlv" }

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) SetParent(parent types.Entity) { organizationSpecificTlv.parent = parent }

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetParent() types.Entity { return organizationSpecificTlv.parent }

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetParentYangName() string { return "last-ccm-received" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv
// Unknown TLVs
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type code of TLV. The type is interface{} with range: 0..255.
    Typecode interface{}

    // Binary data in TLV. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Value interface{}
}

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetFilter() yfilter.YFilter { return unknownTlv.YFilter }

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) SetFilter(yf yfilter.YFilter) { unknownTlv.YFilter = yf }

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetGoName(yname string) string {
    if yname == "typecode" { return "Typecode" }
    if yname == "value" { return "Value" }
    return ""
}

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetSegmentPath() string {
    return "unknown-tlv"
}

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["typecode"] = unknownTlv.Typecode
    leafs["value"] = unknownTlv.Value
    return leafs
}

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetBundleName() string { return "cisco_ios_xr" }

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetYangName() string { return "unknown-tlv" }

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) SetParent(parent types.Entity) { unknownTlv.parent = parent }

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetParent() types.Entity { return unknownTlv.parent }

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetParentYangName() string { return "last-ccm-received" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics
// Peer MEP statistics
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of CCMs received. The type is interface{} with range:
    // 0..18446744073709551615.
    CcmsReceived interface{}

    // Number of CCMs received with an invalid level. The type is interface{} with
    // range: 0..18446744073709551615.
    CcmsWrongLevel interface{}

    // Number of CCMs received with an invalid MAID. The type is interface{} with
    // range: 0..18446744073709551615.
    CcmsInvalidMaid interface{}

    // Number of CCMs received with an invalid interval. The type is interface{}
    // with range: 0..18446744073709551615.
    CcmsInvalidInterval interface{}

    // Number of CCMs received with an invalid source MAC address. The type is
    // interface{} with range: 0..18446744073709551615.
    CcmsInvalidSourceMacAddress interface{}

    // Number of CCMs received with our MEP ID. The type is interface{} with
    // range: 0..18446744073709551615.
    CcmsOurMepId interface{}

    // Number of CCMs received with the Remote Defect Indication bit set. The type
    // is interface{} with range: 0..18446744073709551615.
    CcmsRdi interface{}

    // Number of CCMs received out-of-sequence. The type is interface{} with
    // range: 0..18446744073709551615.
    CcmsOutOfSequence interface{}

    // Sequence number of last CCM received. The type is interface{} with range:
    // 0..4294967295.
    LastCcmSequenceNumber interface{}

    // Elapsed time since last CCM received.
    LastCcmReceivedTime Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime
}

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetGoName(yname string) string {
    if yname == "ccms-received" { return "CcmsReceived" }
    if yname == "ccms-wrong-level" { return "CcmsWrongLevel" }
    if yname == "ccms-invalid-maid" { return "CcmsInvalidMaid" }
    if yname == "ccms-invalid-interval" { return "CcmsInvalidInterval" }
    if yname == "ccms-invalid-source-mac-address" { return "CcmsInvalidSourceMacAddress" }
    if yname == "ccms-our-mep-id" { return "CcmsOurMepId" }
    if yname == "ccms-rdi" { return "CcmsRdi" }
    if yname == "ccms-out-of-sequence" { return "CcmsOutOfSequence" }
    if yname == "last-ccm-sequence-number" { return "LastCcmSequenceNumber" }
    if yname == "last-ccm-received-time" { return "LastCcmReceivedTime" }
    return ""
}

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-ccm-received-time" {
        return &statistics.LastCcmReceivedTime
    }
    return nil
}

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-ccm-received-time"] = &statistics.LastCcmReceivedTime
    return children
}

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccms-received"] = statistics.CcmsReceived
    leafs["ccms-wrong-level"] = statistics.CcmsWrongLevel
    leafs["ccms-invalid-maid"] = statistics.CcmsInvalidMaid
    leafs["ccms-invalid-interval"] = statistics.CcmsInvalidInterval
    leafs["ccms-invalid-source-mac-address"] = statistics.CcmsInvalidSourceMacAddress
    leafs["ccms-our-mep-id"] = statistics.CcmsOurMepId
    leafs["ccms-rdi"] = statistics.CcmsRdi
    leafs["ccms-out-of-sequence"] = statistics.CcmsOutOfSequence
    leafs["last-ccm-sequence-number"] = statistics.LastCcmSequenceNumber
    return leafs
}

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetYangName() string { return "statistics" }

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetParentYangName() string { return "peer-mep" }

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime
// Elapsed time since last CCM received
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetFilter() yfilter.YFilter { return lastCcmReceivedTime.YFilter }

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) SetFilter(yf yfilter.YFilter) { lastCcmReceivedTime.YFilter = yf }

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetSegmentPath() string {
    return "last-ccm-received-time"
}

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = lastCcmReceivedTime.Seconds
    leafs["nanoseconds"] = lastCcmReceivedTime.Nanoseconds
    return leafs
}

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetBundleName() string { return "cisco_ios_xr" }

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetYangName() string { return "last-ccm-received-time" }

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) SetParent(parent types.Entity) { lastCcmReceivedTime.parent = parent }

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetParent() types.Entity { return lastCcmReceivedTime.parent }

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetParentYangName() string { return "statistics" }

