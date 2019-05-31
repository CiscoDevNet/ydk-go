// This module contains general type definitions and identities
// for optical transport models.
package transport_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package transport_types"))
}

type TRIBUTARYPROTOCOLTYPE struct {
}

func (id TRIBUTARYPROTOCOLTYPE) String() string {
	return "openconfig-transport-types:TRIBUTARY_PROTOCOL_TYPE"
}

type PROT1GE struct {
}

func (id PROT1GE) String() string {
	return "openconfig-transport-types:PROT_1GE"
}

type PROTOC48 struct {
}

func (id PROTOC48) String() string {
	return "openconfig-transport-types:PROT_OC48"
}

type PROTSTM16 struct {
}

func (id PROTSTM16) String() string {
	return "openconfig-transport-types:PROT_STM16"
}

type PROT10GELAN struct {
}

func (id PROT10GELAN) String() string {
	return "openconfig-transport-types:PROT_10GE_LAN"
}

type PROT10GEWAN struct {
}

func (id PROT10GEWAN) String() string {
	return "openconfig-transport-types:PROT_10GE_WAN"
}

type PROTOC192 struct {
}

func (id PROTOC192) String() string {
	return "openconfig-transport-types:PROT_OC192"
}

type PROTSTM64 struct {
}

func (id PROTSTM64) String() string {
	return "openconfig-transport-types:PROT_STM64"
}

type PROTOTU2 struct {
}

func (id PROTOTU2) String() string {
	return "openconfig-transport-types:PROT_OTU2"
}

type PROTOTU2E struct {
}

func (id PROTOTU2E) String() string {
	return "openconfig-transport-types:PROT_OTU2E"
}

type PROTOTU1E struct {
}

func (id PROTOTU1E) String() string {
	return "openconfig-transport-types:PROT_OTU1E"
}

type PROTODU2 struct {
}

func (id PROTODU2) String() string {
	return "openconfig-transport-types:PROT_ODU2"
}

type PROTODU2E struct {
}

func (id PROTODU2E) String() string {
	return "openconfig-transport-types:PROT_ODU2E"
}

type PROT40GE struct {
}

func (id PROT40GE) String() string {
	return "openconfig-transport-types:PROT_40GE"
}

type PROTOC768 struct {
}

func (id PROTOC768) String() string {
	return "openconfig-transport-types:PROT_OC768"
}

type PROTSTM256 struct {
}

func (id PROTSTM256) String() string {
	return "openconfig-transport-types:PROT_STM256"
}

type PROTOTU3 struct {
}

func (id PROTOTU3) String() string {
	return "openconfig-transport-types:PROT_OTU3"
}

type PROTODU3 struct {
}

func (id PROTODU3) String() string {
	return "openconfig-transport-types:PROT_ODU3"
}

type PROT100GE struct {
}

func (id PROT100GE) String() string {
	return "openconfig-transport-types:PROT_100GE"
}

type PROT100GMLG struct {
}

func (id PROT100GMLG) String() string {
	return "openconfig-transport-types:PROT_100G_MLG"
}

type PROTOTU4 struct {
}

func (id PROTOTU4) String() string {
	return "openconfig-transport-types:PROT_OTU4"
}

type PROTOTUCN struct {
}

func (id PROTOTUCN) String() string {
	return "openconfig-transport-types:PROT_OTUCN"
}

type PROTODU4 struct {
}

func (id PROTODU4) String() string {
	return "openconfig-transport-types:PROT_ODU4"
}

type TRANSCEIVERFORMFACTORTYPE struct {
}

func (id TRANSCEIVERFORMFACTORTYPE) String() string {
	return "openconfig-transport-types:TRANSCEIVER_FORM_FACTOR_TYPE"
}

type CFP struct {
}

func (id CFP) String() string {
	return "openconfig-transport-types:CFP"
}

type CFP2 struct {
}

func (id CFP2) String() string {
	return "openconfig-transport-types:CFP2"
}

type CFP2ACO struct {
}

func (id CFP2ACO) String() string {
	return "openconfig-transport-types:CFP2_ACO"
}

type CFP4 struct {
}

func (id CFP4) String() string {
	return "openconfig-transport-types:CFP4"
}

type QSFP struct {
}

func (id QSFP) String() string {
	return "openconfig-transport-types:QSFP"
}

type QSFP28 struct {
}

func (id QSFP28) String() string {
	return "openconfig-transport-types:QSFP28"
}

type SFP struct {
}

func (id SFP) String() string {
	return "openconfig-transport-types:SFP"
}

type SFPPLUS struct {
}

func (id SFPPLUS) String() string {
	return "openconfig-transport-types:SFP_PLUS"
}

type XFP struct {
}

func (id XFP) String() string {
	return "openconfig-transport-types:XFP"
}

type X2 struct {
}

func (id X2) String() string {
	return "openconfig-transport-types:X2"
}

type NONPLUGGABLE struct {
}

func (id NONPLUGGABLE) String() string {
	return "openconfig-transport-types:NON_PLUGGABLE"
}

type OTHER struct {
}

func (id OTHER) String() string {
	return "openconfig-transport-types:OTHER"
}

type FIBERCONNECTORTYPE struct {
}

func (id FIBERCONNECTORTYPE) String() string {
	return "openconfig-transport-types:FIBER_CONNECTOR_TYPE"
}

type SCCONNECTOR struct {
}

func (id SCCONNECTOR) String() string {
	return "openconfig-transport-types:SC_CONNECTOR"
}

type LCCONNECTOR struct {
}

func (id LCCONNECTOR) String() string {
	return "openconfig-transport-types:LC_CONNECTOR"
}

type MPOCONNECTOR struct {
}

func (id MPOCONNECTOR) String() string {
	return "openconfig-transport-types:MPO_CONNECTOR"
}

type ETHERNETPMDTYPE struct {
}

func (id ETHERNETPMDTYPE) String() string {
	return "openconfig-transport-types:ETHERNET_PMD_TYPE"
}

type ETH10GBASELRM struct {
}

func (id ETH10GBASELRM) String() string {
	return "openconfig-transport-types:ETH_10GBASE_LRM"
}

type ETH10GBASELR struct {
}

func (id ETH10GBASELR) String() string {
	return "openconfig-transport-types:ETH_10GBASE_LR"
}

type ETH10GBASEZR struct {
}

func (id ETH10GBASEZR) String() string {
	return "openconfig-transport-types:ETH_10GBASE_ZR"
}

type ETH10GBASEER struct {
}

func (id ETH10GBASEER) String() string {
	return "openconfig-transport-types:ETH_10GBASE_ER"
}

type ETH10GBASESR struct {
}

func (id ETH10GBASESR) String() string {
	return "openconfig-transport-types:ETH_10GBASE_SR"
}

type ETH40GBASECR4 struct {
}

func (id ETH40GBASECR4) String() string {
	return "openconfig-transport-types:ETH_40GBASE_CR4"
}

type ETH40GBASESR4 struct {
}

func (id ETH40GBASESR4) String() string {
	return "openconfig-transport-types:ETH_40GBASE_SR4"
}

type ETH40GBASELR4 struct {
}

func (id ETH40GBASELR4) String() string {
	return "openconfig-transport-types:ETH_40GBASE_LR4"
}

type ETH40GBASEER4 struct {
}

func (id ETH40GBASEER4) String() string {
	return "openconfig-transport-types:ETH_40GBASE_ER4"
}

type ETH40GBASEPSM4 struct {
}

func (id ETH40GBASEPSM4) String() string {
	return "openconfig-transport-types:ETH_40GBASE_PSM4"
}

type ETH4X10GBASELR struct {
}

func (id ETH4X10GBASELR) String() string {
	return "openconfig-transport-types:ETH_4X10GBASE_LR"
}

type ETH4X10GBASESR struct {
}

func (id ETH4X10GBASESR) String() string {
	return "openconfig-transport-types:ETH_4X10GBASE_SR"
}

type ETH100GAOC struct {
}

func (id ETH100GAOC) String() string {
	return "openconfig-transport-types:ETH_100G_AOC"
}

type ETH100GACC struct {
}

func (id ETH100GACC) String() string {
	return "openconfig-transport-types:ETH_100G_ACC"
}

type ETH100GBASESR10 struct {
}

func (id ETH100GBASESR10) String() string {
	return "openconfig-transport-types:ETH_100GBASE_SR10"
}

type ETH100GBASESR4 struct {
}

func (id ETH100GBASESR4) String() string {
	return "openconfig-transport-types:ETH_100GBASE_SR4"
}

type ETH100GBASELR4 struct {
}

func (id ETH100GBASELR4) String() string {
	return "openconfig-transport-types:ETH_100GBASE_LR4"
}

type ETH100GBASEER4 struct {
}

func (id ETH100GBASEER4) String() string {
	return "openconfig-transport-types:ETH_100GBASE_ER4"
}

type ETH100GBASECWDM4 struct {
}

func (id ETH100GBASECWDM4) String() string {
	return "openconfig-transport-types:ETH_100GBASE_CWDM4"
}

type ETH100GBASECLR4 struct {
}

func (id ETH100GBASECLR4) String() string {
	return "openconfig-transport-types:ETH_100GBASE_CLR4"
}

type ETH100GBASEPSM4 struct {
}

func (id ETH100GBASEPSM4) String() string {
	return "openconfig-transport-types:ETH_100GBASE_PSM4"
}

type ETH100GBASECR4 struct {
}

func (id ETH100GBASECR4) String() string {
	return "openconfig-transport-types:ETH_100GBASE_CR4"
}

type ETHUNDEFINED struct {
}

func (id ETHUNDEFINED) String() string {
	return "openconfig-transport-types:ETH_UNDEFINED"
}

type SONETAPPLICATIONCODE struct {
}

func (id SONETAPPLICATIONCODE) String() string {
	return "openconfig-transport-types:SONET_APPLICATION_CODE"
}

type VSR20003R2 struct {
}

func (id VSR20003R2) String() string {
	return "openconfig-transport-types:VSR2000_3R2"
}

type VSR20003R3 struct {
}

func (id VSR20003R3) String() string {
	return "openconfig-transport-types:VSR2000_3R3"
}

type VSR20003R5 struct {
}

func (id VSR20003R5) String() string {
	return "openconfig-transport-types:VSR2000_3R5"
}

type SONETUNDEFINED struct {
}

func (id SONETUNDEFINED) String() string {
	return "openconfig-transport-types:SONET_UNDEFINED"
}

type OTNAPPLICATIONCODE struct {
}

func (id OTNAPPLICATIONCODE) String() string {
	return "openconfig-transport-types:OTN_APPLICATION_CODE"
}

type P1L12D1 struct {
}

func (id P1L12D1) String() string {
	return "openconfig-transport-types:P1L1_2D1"
}

type P1S12D2 struct {
}

func (id P1S12D2) String() string {
	return "openconfig-transport-types:P1S1_2D2"
}

type P1L12D2 struct {
}

func (id P1L12D2) String() string {
	return "openconfig-transport-types:P1L1_2D2"
}

type OTNUNDEFINED struct {
}

func (id OTNUNDEFINED) String() string {
	return "openconfig-transport-types:OTN_UNDEFINED"
}

type TRIBUTARYRATECLASSTYPE struct {
}

func (id TRIBUTARYRATECLASSTYPE) String() string {
	return "openconfig-transport-types:TRIBUTARY_RATE_CLASS_TYPE"
}

type TRIBRATE1G struct {
}

func (id TRIBRATE1G) String() string {
	return "openconfig-transport-types:TRIB_RATE_1G"
}

type TRIBRATE2DOT5G struct {
}

func (id TRIBRATE2DOT5G) String() string {
	return "openconfig-transport-types:TRIB_RATE_2.5G"
}

type TRIBRATE10G struct {
}

func (id TRIBRATE10G) String() string {
	return "openconfig-transport-types:TRIB_RATE_10G"
}

type TRIBRATE40G struct {
}

func (id TRIBRATE40G) String() string {
	return "openconfig-transport-types:TRIB_RATE_40G"
}

type TRIBRATE100G struct {
}

func (id TRIBRATE100G) String() string {
	return "openconfig-transport-types:TRIB_RATE_100G"
}

type LOGICALELEMENTPROTOCOLTYPE struct {
}

func (id LOGICALELEMENTPROTOCOLTYPE) String() string {
	return "openconfig-transport-types:LOGICAL_ELEMENT_PROTOCOL_TYPE"
}

type PROTETHERNET struct {
}

func (id PROTETHERNET) String() string {
	return "openconfig-transport-types:PROT_ETHERNET"
}

type PROTOTN struct {
}

func (id PROTOTN) String() string {
	return "openconfig-transport-types:PROT_OTN"
}

type OPTICALCHANNEL struct {
}

func (id OPTICALCHANNEL) String() string {
	return "openconfig-transport-types:OPTICAL_CHANNEL"
}

// AdminStateType represents logical channels in the transponder model.
type AdminStateType string

const (
    // Sets the channel admin state to enabled
    AdminStateType_ENABLED AdminStateType = "ENABLED"

    // Sets the channel admin state to disabled
    AdminStateType_DISABLED AdminStateType = "DISABLED"

    // Sets the channel to maintenance / diagnostic mode
    AdminStateType_MAINT AdminStateType = "MAINT"
)

// LoopbackModeType represents Loopback modes for transponder logical channels
type LoopbackModeType string

const (
    // No loopback is applied
    LoopbackModeType_NONE LoopbackModeType = "NONE"

    // A loopback which directs traffic normally transmitted
    // on the port back to the device as if received on the same
    // port from an external source.
    LoopbackModeType_FACILITY LoopbackModeType = "FACILITY"

    // A loopback which directs traffic received from an external
    // source on the port back out the transmit side of the same
    // port.
    LoopbackModeType_TERMINAL LoopbackModeType = "TERMINAL"
)

