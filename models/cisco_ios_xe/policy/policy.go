// Cisco XE Native Policy Map Yang Model.
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
// All rights reserved.
package policy

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package policy"))
}

// PrecedenceType2
type PrecedenceType2 string

const (
    PrecedenceType2_rsvp PrecedenceType2 = "rsvp"
)

// ClassNameType
type ClassNameType string

const (
    ClassNameType_class_default ClassNameType = "class-default"
)

// PolicePpsBpsType
type PolicePpsBpsType string

const (
    PolicePpsBpsType_pps PolicePpsBpsType = "pps"

    PolicePpsBpsType_bps PolicePpsBpsType = "bps"
)

// PolicePacketsBytesType
type PolicePacketsBytesType string

const (
    PolicePacketsBytesType_packets PolicePacketsBytesType = "packets"

    PolicePacketsBytesType_bytes PolicePacketsBytesType = "bytes"
)

// PolicyActionType
type PolicyActionType string

const (
    PolicyActionType_bandwidth PolicyActionType = "bandwidth"

    PolicyActionType_compression PolicyActionType = "compression"

    PolicyActionType_drop PolicyActionType = "drop"

    PolicyActionType_estimate PolicyActionType = "estimate"

    PolicyActionType_fair_queue PolicyActionType = "fair-queue"

    PolicyActionType_forward PolicyActionType = "forward"

    PolicyActionType_netflow_sampler PolicyActionType = "netflow-sampler"

    PolicyActionType_police PolicyActionType = "police"

    PolicyActionType_priority PolicyActionType = "priority"

    PolicyActionType_queue_limit PolicyActionType = "queue-limit"

    PolicyActionType_random_detect PolicyActionType = "random-detect"

    PolicyActionType_service_policy PolicyActionType = "service-policy"

    PolicyActionType_set PolicyActionType = "set"

    PolicyActionType_shape PolicyActionType = "shape"

    PolicyActionType_trust PolicyActionType = "trust"

    PolicyActionType_dbl PolicyActionType = "dbl"

    PolicyActionType_queue_buffers PolicyActionType = "queue-buffers"
)

// BytesMsUsType
type BytesMsUsType string

const (
    BytesMsUsType_bytes BytesMsUsType = "bytes"

    BytesMsUsType_ms BytesMsUsType = "ms"

    BytesMsUsType_us BytesMsUsType = "us"
)

