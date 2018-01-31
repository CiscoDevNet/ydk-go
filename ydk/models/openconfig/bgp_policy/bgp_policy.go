// This module contains data definitions for BGP routing policy.
// It augments the base routing-policy module with BGP-specific
// options for conditions and actions.
package bgp_policy

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bgp_policy"))
}

// BgpNextHopType represents type definition for specifying next-hop in policy actions
type BgpNextHopType string

const (
    // special designation for local router's own
    // address, i.e., next-hop-self
    BgpNextHopType_SELF BgpNextHopType = "SELF"
)

// BgpSetMedType represents and setting it to the IGP cost (predefined value).
type BgpSetMedType string

const (
    // set the MED value to the IGP cost toward the
    // next hop for the route
    BgpSetMedType_IGP BgpSetMedType = "IGP"
)

// BgpSetCommunityOptionType represents attribute in a policy action
type BgpSetCommunityOptionType string

const (
    // add the specified communities to the existing
    // community attribute
    BgpSetCommunityOptionType_ADD BgpSetCommunityOptionType = "ADD"

    // remove the specified communities from the
    // existing community attribute
    BgpSetCommunityOptionType_REMOVE BgpSetCommunityOptionType = "REMOVE"

    // replace the existing community attribute with
    // the specified communities. If an empty set is
    // specified, this removes the community attribute
    // from the route.
    BgpSetCommunityOptionType_REPLACE BgpSetCommunityOptionType = "REPLACE"
)

