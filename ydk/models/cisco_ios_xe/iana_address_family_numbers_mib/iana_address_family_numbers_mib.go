// The MIB module defines the AddressFamilyNumbers
// textual convention.
package iana_address_family_numbers_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package iana_address_family_numbers_mib"))
}

// AddressFamilyNumbers represents email (iana@iana.org).
type AddressFamilyNumbers string

const (
    AddressFamilyNumbers_other AddressFamilyNumbers = "other"

    AddressFamilyNumbers_ipV4 AddressFamilyNumbers = "ipV4"

    AddressFamilyNumbers_ipV6 AddressFamilyNumbers = "ipV6"

    AddressFamilyNumbers_nsap AddressFamilyNumbers = "nsap"

    AddressFamilyNumbers_hdlc AddressFamilyNumbers = "hdlc"

    AddressFamilyNumbers_bbn1822 AddressFamilyNumbers = "bbn1822"

    AddressFamilyNumbers_all802 AddressFamilyNumbers = "all802"

    AddressFamilyNumbers_e163 AddressFamilyNumbers = "e163"

    AddressFamilyNumbers_e164 AddressFamilyNumbers = "e164"

    AddressFamilyNumbers_f69 AddressFamilyNumbers = "f69"

    AddressFamilyNumbers_x121 AddressFamilyNumbers = "x121"

    AddressFamilyNumbers_ipx AddressFamilyNumbers = "ipx"

    AddressFamilyNumbers_appletalk AddressFamilyNumbers = "appletalk"

    AddressFamilyNumbers_decnetIV AddressFamilyNumbers = "decnetIV"

    AddressFamilyNumbers_banyanVines AddressFamilyNumbers = "banyanVines"

    AddressFamilyNumbers_e164withNsap AddressFamilyNumbers = "e164withNsap"

    AddressFamilyNumbers_dns AddressFamilyNumbers = "dns"

    AddressFamilyNumbers_distinguishedname AddressFamilyNumbers = "distinguishedname"

    AddressFamilyNumbers_asnumber AddressFamilyNumbers = "asnumber"

    AddressFamilyNumbers_xtpoveripv4 AddressFamilyNumbers = "xtpoveripv4"

    AddressFamilyNumbers_xtpoveripv6 AddressFamilyNumbers = "xtpoveripv6"

    AddressFamilyNumbers_xtpnativemodextp AddressFamilyNumbers = "xtpnativemodextp"

    AddressFamilyNumbers_reserved AddressFamilyNumbers = "reserved"
)

