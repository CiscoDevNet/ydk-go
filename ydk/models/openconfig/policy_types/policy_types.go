// This module contains general data definitions for use in routing
// policy.  It can be imported by modules that contain protocol-
// specific policy conditions and actions.
package policy_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package policy_types"))
}

type ATTRIBUTECOMPARISON struct {
}

func (id ATTRIBUTECOMPARISON) String() string {
	return "openconfig-policy-types:ATTRIBUTE_COMPARISON"
}

type ATTRIBUTEEQ struct {
}

func (id ATTRIBUTEEQ) String() string {
	return "openconfig-policy-types:ATTRIBUTE_EQ"
}

type ATTRIBUTEGE struct {
}

func (id ATTRIBUTEGE) String() string {
	return "openconfig-policy-types:ATTRIBUTE_GE"
}

type ATTRIBUTELE struct {
}

func (id ATTRIBUTELE) String() string {
	return "openconfig-policy-types:ATTRIBUTE_LE"
}

type INSTALLPROTOCOLTYPE struct {
}

func (id INSTALLPROTOCOLTYPE) String() string {
	return "openconfig-policy-types:INSTALL_PROTOCOL_TYPE"
}

type BGP struct {
}

func (id BGP) String() string {
	return "openconfig-policy-types:BGP"
}

type ISIS struct {
}

func (id ISIS) String() string {
	return "openconfig-policy-types:ISIS"
}

type OSPF struct {
}

func (id OSPF) String() string {
	return "openconfig-policy-types:OSPF"
}

type OSPF3 struct {
}

func (id OSPF3) String() string {
	return "openconfig-policy-types:OSPF3"
}

type STATIC struct {
}

func (id STATIC) String() string {
	return "openconfig-policy-types:STATIC"
}

type DIRECTLYCONNECTED struct {
}

func (id DIRECTLYCONNECTED) String() string {
	return "openconfig-policy-types:DIRECTLY_CONNECTED"
}

type LOCALAGGREGATE struct {
}

func (id LOCALAGGREGATE) String() string {
	return "openconfig-policy-types:LOCAL_AGGREGATE"
}

// MatchSetOptionsType represents of the members of the defined set
type MatchSetOptionsType string

const (
    // match is true if given value matches any member
    // of the defined set
    MatchSetOptionsType_ANY MatchSetOptionsType = "ANY"

    // match is true if given value matches all
    // members of the defined set
    MatchSetOptionsType_ALL MatchSetOptionsType = "ALL"

    // match is true if given value does not match any
    // member of the defined set
    MatchSetOptionsType_INVERT MatchSetOptionsType = "INVERT"
)

// MatchSetOptionsRestrictedType represents restricted version of the match-set-options-type.
type MatchSetOptionsRestrictedType string

const (
    // match is true if given value matches any member
    // of the defined set
    MatchSetOptionsRestrictedType_ANY MatchSetOptionsRestrictedType = "ANY"

    // match is true if given value does not match any
    // member of the defined set
    MatchSetOptionsRestrictedType_INVERT MatchSetOptionsRestrictedType = "INVERT"
)

