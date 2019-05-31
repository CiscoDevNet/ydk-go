// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_sla_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_sla_datatypes"))
}

// SlaBurstIntervalUnitsEnum represents Sla burst interval units enum
type SlaBurstIntervalUnitsEnum string

const (
    // Send one burst per probe
    SlaBurstIntervalUnitsEnum_once SlaBurstIntervalUnitsEnum = "once"

    // Send bursts within a probe with an interval
    // unit of milliseconds
    SlaBurstIntervalUnitsEnum_milliseconds SlaBurstIntervalUnitsEnum = "milliseconds"

    // Send bursts within a probe with an interval
    // unit of seconds
    SlaBurstIntervalUnitsEnum_seconds SlaBurstIntervalUnitsEnum = "seconds"

    // Send bursts within a probe with an interval
    // unit of minutes
    SlaBurstIntervalUnitsEnum_minutes SlaBurstIntervalUnitsEnum = "minutes"

    // Send bursts within a probe with an interval
    // unit of hours
    SlaBurstIntervalUnitsEnum_hours SlaBurstIntervalUnitsEnum = "hours"
)

// SlaThresholdTypeEnum represents Sla threshold type enum
type SlaThresholdTypeEnum string

const (
    // Stateless threshold
    SlaThresholdTypeEnum_stateless SlaThresholdTypeEnum = "stateless"
)

// SlaSend represents Sla send
type SlaSend string

const (
    // Send individual packets
    SlaSend_packet SlaSend = "packet"

    // Send bursts of packets
    SlaSend_burst SlaSend = "burst"
)

// SlaActionTypeEnum represents Sla action type enum
type SlaActionTypeEnum string

const (
    // Emit a syslog when the threshold is crossed
    SlaActionTypeEnum_log SlaActionTypeEnum = "log"
)

// SlaProbeIntervalDayEnum represents Sla probe interval day enum
type SlaProbeIntervalDayEnum string

const (
    // Schedule every Monday
    SlaProbeIntervalDayEnum_monday SlaProbeIntervalDayEnum = "monday"

    // Schedule every Tuesday
    SlaProbeIntervalDayEnum_tuesday SlaProbeIntervalDayEnum = "tuesday"

    // Schedule every Wednesday
    SlaProbeIntervalDayEnum_wednesday SlaProbeIntervalDayEnum = "wednesday"

    // Schedule every Thursday
    SlaProbeIntervalDayEnum_thursday SlaProbeIntervalDayEnum = "thursday"

    // Schedule every Friday
    SlaProbeIntervalDayEnum_friday SlaProbeIntervalDayEnum = "friday"

    // Schedule every Saturday
    SlaProbeIntervalDayEnum_saturday SlaProbeIntervalDayEnum = "saturday"

    // Schedule every Sunday
    SlaProbeIntervalDayEnum_sunday SlaProbeIntervalDayEnum = "sunday"
)

// SlaPacketIntervalUnitsEnum represents Sla packet interval units enum
type SlaPacketIntervalUnitsEnum string

const (
    // Send one packet per burst
    SlaPacketIntervalUnitsEnum_once SlaPacketIntervalUnitsEnum = "once"

    // Send packets with an interval unit of
    // milliseconds
    SlaPacketIntervalUnitsEnum_milliseconds SlaPacketIntervalUnitsEnum = "milliseconds"

    // Send packets with an interval unit of seconds
    SlaPacketIntervalUnitsEnum_seconds SlaPacketIntervalUnitsEnum = "seconds"

    // Send packets with an interval unit of minutes
    SlaPacketIntervalUnitsEnum_minutes SlaPacketIntervalUnitsEnum = "minutes"

    // Send packets with an interval unit of hours
    SlaPacketIntervalUnitsEnum_hours SlaPacketIntervalUnitsEnum = "hours"
)

// SlaOnDemandRepeatIntervalUnitsEnum represents Sla on demand repeat interval units enum
type SlaOnDemandRepeatIntervalUnitsEnum string

const (
    // Schedule probes to repeat with an interval unit
    // of seconds
    SlaOnDemandRepeatIntervalUnitsEnum_seconds SlaOnDemandRepeatIntervalUnitsEnum = "seconds"

    // Schedule probes to repeat with an interval unit
    // of minutes
    SlaOnDemandRepeatIntervalUnitsEnum_minutes SlaOnDemandRepeatIntervalUnitsEnum = "minutes"

    // Schedule probes to repeat with an interval unit
    // of hours
    SlaOnDemandRepeatIntervalUnitsEnum_hours SlaOnDemandRepeatIntervalUnitsEnum = "hours"
)

// SlaProbeDurationUnitsEnum represents Sla probe duration units enum
type SlaProbeDurationUnitsEnum string

const (
    // Schedule probes to run with a duration unit of
    // seconds
    SlaProbeDurationUnitsEnum_seconds SlaProbeDurationUnitsEnum = "seconds"

    // Schedule probes to run with a duration unit of
    // minutes
    SlaProbeDurationUnitsEnum_minutes SlaProbeDurationUnitsEnum = "minutes"

    // Schedule probes to run with a duration unit of
    // hours
    SlaProbeDurationUnitsEnum_hours SlaProbeDurationUnitsEnum = "hours"

    // Schedule probes to run for a duration of 1 day
    SlaProbeDurationUnitsEnum_day SlaProbeDurationUnitsEnum = "day"

    // Schedule probes to run for a duration of 1 week
    SlaProbeDurationUnitsEnum_week SlaProbeDurationUnitsEnum = "week"
)

// SlaStatisticTypeEnum represents Sla statistic type enum
type SlaStatisticTypeEnum string

const (
    // Collect round trip delay metric data
    SlaStatisticTypeEnum_round_trip_delay SlaStatisticTypeEnum = "round-trip-delay"

    // Collect one way delay source->dest metric data
    SlaStatisticTypeEnum_one_way_delay_sd SlaStatisticTypeEnum = "one-way-delay-sd"

    // Collect one way delay dest->source metric data
    SlaStatisticTypeEnum_one_way_delay_ds SlaStatisticTypeEnum = "one-way-delay-ds"

    // Collect round trip delay metric data
    SlaStatisticTypeEnum_round_trip_jitter SlaStatisticTypeEnum = "round-trip-jitter"

    // Collect one way jitter source->dest metric data
    SlaStatisticTypeEnum_one_way_jitter_sd SlaStatisticTypeEnum = "one-way-jitter-sd"

    // Collect one way jitter dest->source metric data
    SlaStatisticTypeEnum_one_way_jitter_ds SlaStatisticTypeEnum = "one-way-jitter-ds"

    // Collect one way loss source->dest metric data
    SlaStatisticTypeEnum_one_way_loss_sd SlaStatisticTypeEnum = "one-way-loss-sd"

    // Collect one way loss dest->source metric data
    SlaStatisticTypeEnum_one_way_loss_ds SlaStatisticTypeEnum = "one-way-loss-ds"
)

// SlaBucketsSizeUnitsEnum represents Sla buckets size units enum
type SlaBucketsSizeUnitsEnum string

const (
    // Store results as a number of buckets per probe
    // - note that this option has been DEPRECATED
    SlaBucketsSizeUnitsEnum_buckets_per_probe SlaBucketsSizeUnitsEnum = "buckets-per-probe"

    // Store results as a number of probes per bucket
    SlaBucketsSizeUnitsEnum_probes_per_bucket SlaBucketsSizeUnitsEnum = "probes-per-bucket"
)

// SlaThresholdConditionEnum represents Sla threshold condition enum
type SlaThresholdConditionEnum string

const (
    // Threshold is breached when the maximum value
    // crosses the configured threshold value
    SlaThresholdConditionEnum_max SlaThresholdConditionEnum = "max"

    // Threshold is breached when the mean value
    // crosses the configured threshold value
    SlaThresholdConditionEnum_mean SlaThresholdConditionEnum = "mean"

    // Threshold is breached when the sample count in
    // bins in and above a certain bin number crosses
    // the configured sample count
    SlaThresholdConditionEnum_sample_count SlaThresholdConditionEnum = "sample-count"
)

// SlaProbeIntervalUnitsEnum represents Sla probe interval units enum
type SlaProbeIntervalUnitsEnum string

const (
    // Schedule probes to run with an interval unit of
    // minutes
    SlaProbeIntervalUnitsEnum_minutes SlaProbeIntervalUnitsEnum = "minutes"

    // Schedule probes to run with an interval unit of
    // hours
    SlaProbeIntervalUnitsEnum_hours SlaProbeIntervalUnitsEnum = "hours"

    // Schedule probes to run every day
    SlaProbeIntervalUnitsEnum_day SlaProbeIntervalUnitsEnum = "day"

    // Schedule probes to run every week
    SlaProbeIntervalUnitsEnum_week SlaProbeIntervalUnitsEnum = "week"
)

// SlaOnDemandStartMonthEnum represents Sla on demand start month enum
type SlaOnDemandStartMonthEnum string

const (
    // January
    SlaOnDemandStartMonthEnum_january SlaOnDemandStartMonthEnum = "january"

    // February
    SlaOnDemandStartMonthEnum_february SlaOnDemandStartMonthEnum = "february"

    // March
    SlaOnDemandStartMonthEnum_march SlaOnDemandStartMonthEnum = "march"

    // April
    SlaOnDemandStartMonthEnum_april SlaOnDemandStartMonthEnum = "april"

    // May
    SlaOnDemandStartMonthEnum_may SlaOnDemandStartMonthEnum = "may"

    // June
    SlaOnDemandStartMonthEnum_june SlaOnDemandStartMonthEnum = "june"

    // July
    SlaOnDemandStartMonthEnum_july SlaOnDemandStartMonthEnum = "july"

    // August
    SlaOnDemandStartMonthEnum_august SlaOnDemandStartMonthEnum = "august"

    // September
    SlaOnDemandStartMonthEnum_september SlaOnDemandStartMonthEnum = "september"

    // October
    SlaOnDemandStartMonthEnum_october SlaOnDemandStartMonthEnum = "october"

    // November
    SlaOnDemandStartMonthEnum_november SlaOnDemandStartMonthEnum = "november"

    // December
    SlaOnDemandStartMonthEnum_december SlaOnDemandStartMonthEnum = "december"
)

// SlaPaddingPattern represents Sla padding pattern
type SlaPaddingPattern string

const (
    // Use an optionally specified hex pattern for
    // packet padding
    SlaPaddingPattern_hex SlaPaddingPattern = "hex"

    // Use a pseudo-random bit sequence for packet
    // padding
    SlaPaddingPattern_pseudo_random SlaPaddingPattern = "pseudo-random"
)

// SlaOnDemandStartTimeTypesEnum represents Sla on demand start time types enum
type SlaOnDemandStartTimeTypesEnum string

const (
    // Start immediately
    SlaOnDemandStartTimeTypesEnum_now SlaOnDemandStartTimeTypesEnum = "now"

    // Start at a specified time
    SlaOnDemandStartTimeTypesEnum_absolute SlaOnDemandStartTimeTypesEnum = "absolute"

    // Start after a specified period
    SlaOnDemandStartTimeTypesEnum_relative SlaOnDemandStartTimeTypesEnum = "relative"
)

// SlaOnDemandProbeDurationUnitsEnum represents Sla on demand probe duration units enum
type SlaOnDemandProbeDurationUnitsEnum string

const (
    // Schedule probes to run with a duration unit of
    // seconds
    SlaOnDemandProbeDurationUnitsEnum_seconds SlaOnDemandProbeDurationUnitsEnum = "seconds"

    // Schedule probes to run with a duration unit of
    // minutes
    SlaOnDemandProbeDurationUnitsEnum_minutes SlaOnDemandProbeDurationUnitsEnum = "minutes"

    // Schedule probes to run with a duration unit of
    // hours
    SlaOnDemandProbeDurationUnitsEnum_hours SlaOnDemandProbeDurationUnitsEnum = "hours"
)

// SlaOnDemandStartTimeRelativeUnitsEnum represents Sla on demand start time relative units enum
type SlaOnDemandStartTimeRelativeUnitsEnum string

const (
    // Schedule probe to start after a unit of seconds
    SlaOnDemandStartTimeRelativeUnitsEnum_seconds SlaOnDemandStartTimeRelativeUnitsEnum = "seconds"

    // Schedule probe to start after a unit of minutes
    SlaOnDemandStartTimeRelativeUnitsEnum_minutes SlaOnDemandStartTimeRelativeUnitsEnum = "minutes"

    // Schedule probe to start after a unit of hours
    SlaOnDemandStartTimeRelativeUnitsEnum_hours SlaOnDemandStartTimeRelativeUnitsEnum = "hours"
)

