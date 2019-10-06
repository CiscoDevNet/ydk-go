// This module contains general data definitions for use in ISIS YANG
// model.
package isis_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package isis_types"))
}

type OVERLOADRESETTRIGGERTYPE struct {
}

func (id OVERLOADRESETTRIGGERTYPE) String() string {
	return "openconfig-isis-types:OVERLOAD_RESET_TRIGGER_TYPE"
}

type WAITFORSYSTEM struct {
}

func (id WAITFORSYSTEM) String() string {
	return "openconfig-isis-types:WAIT_FOR_SYSTEM"
}

type IPV6UNICAST struct {
}

func (id IPV6UNICAST) String() string {
	return "openconfig-isis-types:IPV6_UNICAST"
}

type SAFITYPE struct {
}

func (id SAFITYPE) String() string {
	return "openconfig-isis-types:SAFI_TYPE"
}

type UNICAST struct {
}

func (id UNICAST) String() string {
	return "openconfig-isis-types:UNICAST"
}

type MTTYPE struct {
}

func (id MTTYPE) String() string {
	return "openconfig-isis-types:MT_TYPE"
}

type MULTICAST struct {
}

func (id MULTICAST) String() string {
	return "openconfig-isis-types:MULTICAST"
}

type AFISAFITYPE struct {
}

func (id AFISAFITYPE) String() string {
	return "openconfig-isis-types:AFI_SAFI_TYPE"
}

type IPV4 struct {
}

func (id IPV4) String() string {
	return "openconfig-isis-types:IPV4"
}

type IPV6 struct {
}

func (id IPV6) String() string {
	return "openconfig-isis-types:IPV6"
}

type IPV4MULTICAST struct {
}

func (id IPV4MULTICAST) String() string {
	return "openconfig-isis-types:IPV4_MULTICAST"
}

type WAITFORBGP struct {
}

func (id WAITFORBGP) String() string {
	return "openconfig-isis-types:WAIT_FOR_BGP"
}

type IPV6MULTICAST struct {
}

func (id IPV6MULTICAST) String() string {
	return "openconfig-isis-types:IPV6_MULTICAST"
}

type AFITYPE struct {
}

func (id AFITYPE) String() string {
	return "openconfig-isis-types:AFI_TYPE"
}

type IPV4UNICAST struct {
}

func (id IPV4UNICAST) String() string {
	return "openconfig-isis-types:IPV4_UNICAST"
}

// IsisInterfaceAdjState represents This type defines the state of the interface.
type IsisInterfaceAdjState string

const (
    // This state describes that adjacency is established.
    IsisInterfaceAdjState_UP IsisInterfaceAdjState = "UP"

    // This state describes that adjacency is NOT established.
    IsisInterfaceAdjState_DOWN IsisInterfaceAdjState = "DOWN"

    // This state describes that adjacency is establishing.
    IsisInterfaceAdjState_INIT IsisInterfaceAdjState = "INIT"

    // This state describes that adjacency is failed.
    IsisInterfaceAdjState_FAILED IsisInterfaceAdjState = "FAILED"
)

// LevelType represents This type defines ISIS level types
type LevelType string

const (
    // This enum describes ISIS level 1
    LevelType_LEVEL_1 LevelType = "LEVEL_1"

    // This enum describes ISIS level 2
    LevelType_LEVEL_2 LevelType = "LEVEL_2"

    // This enum describes ISIS level 1-2
    LevelType_LEVEL_1_2 LevelType = "LEVEL_1_2"
)

// MetricType represents This type defines ISIS metric type
type MetricType string

const (
    // This enum describes internal route type
    MetricType_INTERNAL MetricType = "INTERNAL"

    // This enum describes external route type
    MetricType_EXTERNAL MetricType = "EXTERNAL"
)

// HelloPaddingType represents This type defines ISIS hello padding type
type HelloPaddingType string

const (
    // This enum describes strict padding
    HelloPaddingType_STRICT HelloPaddingType = "STRICT"

    // This enum describes loose padding
    HelloPaddingType_LOOSE HelloPaddingType = "LOOSE"

    // This enum describes adaptive padding
    HelloPaddingType_ADAPTIVE HelloPaddingType = "ADAPTIVE"

    // This enum disables padding
    HelloPaddingType_DISABLE HelloPaddingType = "DISABLE"
)

// AdaptiveTimerType represents This type defines ISIS adaptive timer types
type AdaptiveTimerType string

const (
    // This enum describes linear algorithm timer
    AdaptiveTimerType_LINEAR AdaptiveTimerType = "LINEAR"

    // This enum describes exponential algorithm timer
    AdaptiveTimerType_EXPONENTIAL AdaptiveTimerType = "EXPONENTIAL"
)

// CircuitType represents This type defines ISIS interface types 
type CircuitType string

const (
    // This enum describes a point-to-point interface
    CircuitType_POINT_TO_POINT CircuitType = "POINT_TO_POINT"

    // This enum describes a broadcast interface
    CircuitType_BROADCAST CircuitType = "BROADCAST"
)

// MetricStyle represents This type defines ISIS metric styles
type MetricStyle string

const (
    // This enum describes narrow metric style
    MetricStyle_NARROW_METRIC MetricStyle = "NARROW_METRIC"

    // This enum describes wide metric style
    MetricStyle_WIDE_METRIC MetricStyle = "WIDE_METRIC"
)

