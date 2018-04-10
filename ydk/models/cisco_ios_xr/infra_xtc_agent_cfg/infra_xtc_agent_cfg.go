// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-xtc-agent package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-segment-routing-ms-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_xtc_agent_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_xtc_agent_cfg"))
}

// XtcMetricValue represents Xtc metric value
type XtcMetricValue string

const (
    // Relative metric value
    XtcMetricValue_relative XtcMetricValue = "relative"

    // Absolute metric value
    XtcMetricValue_absolute XtcMetricValue = "absolute"
)

// XtcBindingSidexplicitRule represents Xtc binding sidexplicit rule
type XtcBindingSidexplicitRule string

const (
    // Fallback dynamic option
    XtcBindingSidexplicitRule_fallback_dynamic XtcBindingSidexplicitRule = "fallback-dynamic"

    // SRLB enforcement option
    XtcBindingSidexplicitRule_enforce_srlb XtcBindingSidexplicitRule = "enforce-srlb"
)

// XtcPathHop represents Xtc path hop
type XtcPathHop string

const (
    // Segment-routing MPLS
    XtcPathHop_mpls XtcPathHop = "mpls"

    // Segment-routing v6
    XtcPathHop_srv6 XtcPathHop = "srv6"
)

// XtcAffinityRule represents Xtc affinity rule
type XtcAffinityRule string

const (
    // Include-all affinity rule
    XtcAffinityRule_affinity_include_all XtcAffinityRule = "affinity-include-all"

    // Exclude-any affinity rule
    XtcAffinityRule_affinity_exclude_any XtcAffinityRule = "affinity-exclude-any"

    // Include-any affinity rule
    XtcAffinityRule_affinity_include_any XtcAffinityRule = "affinity-include-any"
)

// XtcPath represents Xtc path
type XtcPath string

const (
    // Explicit
    XtcPath_explicit XtcPath = "explicit"

    // Dynamic
    XtcPath_dynamic XtcPath = "dynamic"
)

// XtcMetric represents Xtc metric
type XtcMetric string

const (
    // IGP metric type
    XtcMetric_igp XtcMetric = "igp"

    // TE metric type
    XtcMetric_te XtcMetric = "te"
)

// XtcSegment represents Xtc segment
type XtcSegment string

const (
    // IPv4 Address
    XtcSegment_ipv4_address XtcSegment = "ipv4-address"

    // MPLS Label
    XtcSegment_mpls_label XtcSegment = "mpls-label"
)

// XtcBindingSid represents Xtc binding sid
type XtcBindingSid string

const (
    // Use specified BSID MPLS label
    XtcBindingSid_mpls_label_specified XtcBindingSid = "mpls-label-specified"

    // Allocate BSID MPLS label
    XtcBindingSid_mpls_label_any XtcBindingSid = "mpls-label-any"
)

// XtcDisjointness represents Xtc disjointness
type XtcDisjointness string

const (
    // Link Disjointness
    XtcDisjointness_link XtcDisjointness = "link"

    // Node Disjointness
    XtcDisjointness_node XtcDisjointness = "node"

    // SRLG Disjointness
    XtcDisjointness_srlg XtcDisjointness = "srlg"

    // SRLG Node Disjointness
    XtcDisjointness_srlg_node XtcDisjointness = "srlg-node"
)

// XtcEndPoint represents Xtc end point
type XtcEndPoint string

const (
    // IPv4 endpoint address
    XtcEndPoint_end_point_type_ipv4 XtcEndPoint = "end-point-type-ipv4"

    // IPv6 endpoint address
    XtcEndPoint_end_point_type_ipv6 XtcEndPoint = "end-point-type-ipv6"
)

