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

type Atmnotrafficdescriptor struct {
}

func (id Atmnotrafficdescriptor) String() string {
	return "ATM-TC-MIB:atmNoTrafficDescriptor"
}

type Atmnoclpnoscr struct {
}

func (id Atmnoclpnoscr) String() string {
	return "ATM-TC-MIB:atmNoClpNoScr"
}

type Atmclpnotaggingnoscr struct {
}

func (id Atmclpnotaggingnoscr) String() string {
	return "ATM-TC-MIB:atmClpNoTaggingNoScr"
}

type Atmclptaggingnoscr struct {
}

func (id Atmclptaggingnoscr) String() string {
	return "ATM-TC-MIB:atmClpTaggingNoScr"
}

type Atmnoclpscr struct {
}

func (id Atmnoclpscr) String() string {
	return "ATM-TC-MIB:atmNoClpScr"
}

type Atmclpnotaggingscr struct {
}

func (id Atmclpnotaggingscr) String() string {
	return "ATM-TC-MIB:atmClpNoTaggingScr"
}

type Atmclptaggingscr struct {
}

func (id Atmclptaggingscr) String() string {
	return "ATM-TC-MIB:atmClpTaggingScr"
}

type Atmclpnotaggingmcr struct {
}

func (id Atmclpnotaggingmcr) String() string {
	return "ATM-TC-MIB:atmClpNoTaggingMcr"
}

type Atmclptransparentnoscr struct {
}

func (id Atmclptransparentnoscr) String() string {
	return "ATM-TC-MIB:atmClpTransparentNoScr"
}

type Atmclptransparentscr struct {
}

func (id Atmclptransparentscr) String() string {
	return "ATM-TC-MIB:atmClpTransparentScr"
}

type Atmnoclptaggingnoscr struct {
}

func (id Atmnoclptaggingnoscr) String() string {
	return "ATM-TC-MIB:atmNoClpTaggingNoScr"
}

type Atmnoclpnoscrcdvt struct {
}

func (id Atmnoclpnoscrcdvt) String() string {
	return "ATM-TC-MIB:atmNoClpNoScrCdvt"
}

type Atmnoclpscrcdvt struct {
}

func (id Atmnoclpscrcdvt) String() string {
	return "ATM-TC-MIB:atmNoClpScrCdvt"
}

type Atmclpnotaggingscrcdvt struct {
}

func (id Atmclpnotaggingscrcdvt) String() string {
	return "ATM-TC-MIB:atmClpNoTaggingScrCdvt"
}

type Atmclptaggingscrcdvt struct {
}

func (id Atmclptaggingscrcdvt) String() string {
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

