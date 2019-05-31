// This MIB Module provides Textual Conventions
// and OBJECT-IDENTITY Objects to be used by
// ATM systems.
package atm_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package atm_tc_mib"))
}

type AtmNoTrafficDescriptor struct {
}

func (id AtmNoTrafficDescriptor) String() string {
	return "ATM-TC-MIB:atmNoTrafficDescriptor"
}

type AtmNoClpNoScr struct {
}

func (id AtmNoClpNoScr) String() string {
	return "ATM-TC-MIB:atmNoClpNoScr"
}

type AtmClpNoTaggingNoScr struct {
}

func (id AtmClpNoTaggingNoScr) String() string {
	return "ATM-TC-MIB:atmClpNoTaggingNoScr"
}

type AtmClpTaggingNoScr struct {
}

func (id AtmClpTaggingNoScr) String() string {
	return "ATM-TC-MIB:atmClpTaggingNoScr"
}

type AtmNoClpScr struct {
}

func (id AtmNoClpScr) String() string {
	return "ATM-TC-MIB:atmNoClpScr"
}

type AtmClpNoTaggingScr struct {
}

func (id AtmClpNoTaggingScr) String() string {
	return "ATM-TC-MIB:atmClpNoTaggingScr"
}

type AtmClpTaggingScr struct {
}

func (id AtmClpTaggingScr) String() string {
	return "ATM-TC-MIB:atmClpTaggingScr"
}

type AtmClpNoTaggingMcr struct {
}

func (id AtmClpNoTaggingMcr) String() string {
	return "ATM-TC-MIB:atmClpNoTaggingMcr"
}

type AtmClpTransparentNoScr struct {
}

func (id AtmClpTransparentNoScr) String() string {
	return "ATM-TC-MIB:atmClpTransparentNoScr"
}

type AtmClpTransparentScr struct {
}

func (id AtmClpTransparentScr) String() string {
	return "ATM-TC-MIB:atmClpTransparentScr"
}

type AtmNoClpTaggingNoScr struct {
}

func (id AtmNoClpTaggingNoScr) String() string {
	return "ATM-TC-MIB:atmNoClpTaggingNoScr"
}

type AtmNoClpNoScrCdvt struct {
}

func (id AtmNoClpNoScrCdvt) String() string {
	return "ATM-TC-MIB:atmNoClpNoScrCdvt"
}

type AtmNoClpScrCdvt struct {
}

func (id AtmNoClpScrCdvt) String() string {
	return "ATM-TC-MIB:atmNoClpScrCdvt"
}

type AtmClpNoTaggingScrCdvt struct {
}

func (id AtmClpNoTaggingScrCdvt) String() string {
	return "ATM-TC-MIB:atmClpNoTaggingScrCdvt"
}

type AtmClpTaggingScrCdvt struct {
}

func (id AtmClpTaggingScrCdvt) String() string {
	return "ATM-TC-MIB:atmClpTaggingScrCdvt"
}

// AtmConnCastType represents   are to the leaf of the p2mp connection.
type AtmConnCastType string

const (
    AtmConnCastType_p2p AtmConnCastType = "p2p"

    AtmConnCastType_p2mpRoot AtmConnCastType = "p2mpRoot"

    AtmConnCastType_p2mpLeaf AtmConnCastType = "p2mpLeaf"
)

// AtmConnKind represents cross-connected to an svcIncoming or an spvcInitiator.
type AtmConnKind string

const (
    AtmConnKind_pvc AtmConnKind = "pvc"

    AtmConnKind_svcIncoming AtmConnKind = "svcIncoming"

    AtmConnKind_svcOutgoing AtmConnKind = "svcOutgoing"

    AtmConnKind_spvcInitiator AtmConnKind = "spvcInitiator"

    AtmConnKind_spvcTarget AtmConnKind = "spvcTarget"
)

// AtmInterfaceType represents    signalling disabled.
type AtmInterfaceType string

const (
    AtmInterfaceType_other AtmInterfaceType = "other"

    AtmInterfaceType_autoConfig AtmInterfaceType = "autoConfig"

    AtmInterfaceType_ituDss2 AtmInterfaceType = "ituDss2"

    AtmInterfaceType_atmfUni3Dot0 AtmInterfaceType = "atmfUni3Dot0"

    AtmInterfaceType_atmfUni3Dot1 AtmInterfaceType = "atmfUni3Dot1"

    AtmInterfaceType_atmfUni4Dot0 AtmInterfaceType = "atmfUni4Dot0"

    AtmInterfaceType_atmfIispUni3Dot0 AtmInterfaceType = "atmfIispUni3Dot0"

    AtmInterfaceType_atmfIispUni3Dot1 AtmInterfaceType = "atmfIispUni3Dot1"

    AtmInterfaceType_atmfIispUni4Dot0 AtmInterfaceType = "atmfIispUni4Dot0"

    AtmInterfaceType_atmfPnni1Dot0 AtmInterfaceType = "atmfPnni1Dot0"

    AtmInterfaceType_atmfBici2Dot0 AtmInterfaceType = "atmfBici2Dot0"

    AtmInterfaceType_atmfUniPvcOnly AtmInterfaceType = "atmfUniPvcOnly"

    AtmInterfaceType_atmfNniPvcOnly AtmInterfaceType = "atmfNniPvcOnly"
)

// AtmServiceCategory represents The service category for a connection.
type AtmServiceCategory string

const (
    AtmServiceCategory_other AtmServiceCategory = "other"

    AtmServiceCategory_cbr AtmServiceCategory = "cbr"

    AtmServiceCategory_rtVbr AtmServiceCategory = "rtVbr"

    AtmServiceCategory_nrtVbr AtmServiceCategory = "nrtVbr"

    AtmServiceCategory_abr AtmServiceCategory = "abr"

    AtmServiceCategory_ubr AtmServiceCategory = "ubr"
)

// AtmVorXAdminStatus represents link or cross-connect.
type AtmVorXAdminStatus string

const (
    AtmVorXAdminStatus_up AtmVorXAdminStatus = "up"

    AtmVorXAdminStatus_down AtmVorXAdminStatus = "down"
)

// AtmVorXOperStatus represents interface(s) is down or unknown respectively.
type AtmVorXOperStatus string

const (
    AtmVorXOperStatus_up AtmVorXOperStatus = "up"

    AtmVorXOperStatus_down AtmVorXOperStatus = "down"

    AtmVorXOperStatus_unknown AtmVorXOperStatus = "unknown"
)

