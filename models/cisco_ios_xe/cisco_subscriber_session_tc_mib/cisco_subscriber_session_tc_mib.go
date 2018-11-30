// This MIB module defines textual conventions describing
// subscriber sessions.
package cisco_subscriber_session_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_subscriber_session_tc_mib"))
}

// SubSessionType represents indicate this fact.
type SubSessionType string

const (
    SubSessionType_all SubSessionType = "all"

    SubSessionType_other SubSessionType = "other"

    SubSessionType_pppSubscriber SubSessionType = "pppSubscriber"

    SubSessionType_pppoeSubscriber SubSessionType = "pppoeSubscriber"

    SubSessionType_l2tpSubscriber SubSessionType = "l2tpSubscriber"

    SubSessionType_l2fSubscriber SubSessionType = "l2fSubscriber"

    SubSessionType_ipInterfaceSubscriber SubSessionType = "ipInterfaceSubscriber"

    SubSessionType_ipPktSubscriber SubSessionType = "ipPktSubscriber"

    SubSessionType_ipDhcpv4Subscriber SubSessionType = "ipDhcpv4Subscriber"

    SubSessionType_ipRadiusSubscriber SubSessionType = "ipRadiusSubscriber"

    SubSessionType_l2MacSubscriber SubSessionType = "l2MacSubscriber"

    SubSessionType_l2Dhcpv4Subscriber SubSessionType = "l2Dhcpv4Subscriber"

    SubSessionType_l2RadiusSubscriber SubSessionType = "l2RadiusSubscriber"
)

// SubSessionRedundancyMode represents         a fail-over event (e.g., linecard failure).
type SubSessionRedundancyMode string

const (
    SubSessionRedundancyMode_none SubSessionRedundancyMode = "none"

    SubSessionRedundancyMode_other SubSessionRedundancyMode = "other"

    SubSessionRedundancyMode_active SubSessionRedundancyMode = "active"

    SubSessionRedundancyMode_standby SubSessionRedundancyMode = "standby"
)

// SubSessionState represents         system has established the subscriber session.
type SubSessionState string

const (
    SubSessionState_other SubSessionState = "other"

    SubSessionState_pending SubSessionState = "pending"

    SubSessionState_up SubSessionState = "up"
)

