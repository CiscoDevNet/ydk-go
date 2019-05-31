// This MIB module defines textual conventions describing
// subscriber session identities.  A subscriber session identity
// consists of data associated with a subscriber session serving as
// credentials used to determine authority, status, rights, or
// entitlement to privileges.
package cisco_subscriber_identity_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_subscriber_identity_tc_mib"))
}

// SubSessionIdentity represents          session.
type SubSessionIdentity string

const (
    SubSessionIdentity_other SubSessionIdentity = "other"

    SubSessionIdentity_ifIndex SubSessionIdentity = "ifIndex"

    SubSessionIdentity_subscriberLabel SubSessionIdentity = "subscriberLabel"

    SubSessionIdentity_macAddress SubSessionIdentity = "macAddress"

    SubSessionIdentity_nativeVrf SubSessionIdentity = "nativeVrf"

    SubSessionIdentity_nativeIpAddress SubSessionIdentity = "nativeIpAddress"

    SubSessionIdentity_domainVrf SubSessionIdentity = "domainVrf"

    SubSessionIdentity_domainIpAddress SubSessionIdentity = "domainIpAddress"

    SubSessionIdentity_pbhk SubSessionIdentity = "pbhk"

    SubSessionIdentity_remoteId SubSessionIdentity = "remoteId"

    SubSessionIdentity_circuitId SubSessionIdentity = "circuitId"

    SubSessionIdentity_nasPort SubSessionIdentity = "nasPort"

    SubSessionIdentity_domain SubSessionIdentity = "domain"

    SubSessionIdentity_username SubSessionIdentity = "username"

    SubSessionIdentity_acctSessionId SubSessionIdentity = "acctSessionId"

    SubSessionIdentity_dnis SubSessionIdentity = "dnis"

    SubSessionIdentity_media SubSessionIdentity = "media"

    SubSessionIdentity_mlpNegotiated SubSessionIdentity = "mlpNegotiated"

    SubSessionIdentity_protocol SubSessionIdentity = "protocol"

    SubSessionIdentity_serviceName SubSessionIdentity = "serviceName"

    SubSessionIdentity_dhcpClass SubSessionIdentity = "dhcpClass"

    SubSessionIdentity_tunnelName SubSessionIdentity = "tunnelName"
)

// SubscriberMediaType represents object using this textual convention is not valid.
type SubscriberMediaType string

const (
    SubscriberMediaType_other SubscriberMediaType = "other"

    SubscriberMediaType_async SubscriberMediaType = "async"

    SubscriberMediaType_atm SubscriberMediaType = "atm"

    SubscriberMediaType_ethernet SubscriberMediaType = "ethernet"

    SubscriberMediaType_ip SubscriberMediaType = "ip"

    SubscriberMediaType_isdn SubscriberMediaType = "isdn"

    SubscriberMediaType_mpls SubscriberMediaType = "mpls"

    SubscriberMediaType_sync SubscriberMediaType = "sync"
)

// SubscriberProtocolType represents object using this textual convention is not valid.
type SubscriberProtocolType string

const (
    SubscriberProtocolType_other SubscriberProtocolType = "other"

    SubscriberProtocolType_atom SubscriberProtocolType = "atom"

    SubscriberProtocolType_ip SubscriberProtocolType = "ip"

    SubscriberProtocolType_psdn SubscriberProtocolType = "psdn"

    SubscriberProtocolType_ppp SubscriberProtocolType = "ppp"

    SubscriberProtocolType_vpdn SubscriberProtocolType = "vpdn"
)

