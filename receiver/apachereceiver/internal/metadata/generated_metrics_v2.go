// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"fmt"
	"strconv"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for apachereceiver metrics.
type MetricsSettings struct {
	ApacheCurrentConnections MetricSettings `mapstructure:"apache.current_connections"`
	ApacheRequests           MetricSettings `mapstructure:"apache.requests"`
	ApacheScoreboard         MetricSettings `mapstructure:"apache.scoreboard"`
	ApacheTraffic            MetricSettings `mapstructure:"apache.traffic"`
	ApacheUptime             MetricSettings `mapstructure:"apache.uptime"`
	ApacheWorkers            MetricSettings `mapstructure:"apache.workers"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		ApacheCurrentConnections: MetricSettings{
			Enabled: true,
		},
		ApacheRequests: MetricSettings{
			Enabled: true,
		},
		ApacheScoreboard: MetricSettings{
			Enabled: true,
		},
		ApacheTraffic: MetricSettings{
			Enabled: true,
		},
		ApacheUptime: MetricSettings{
			Enabled: true,
		},
		ApacheWorkers: MetricSettings{
			Enabled: true,
		},
	}
}

// AttributeScoreboardState specifies the a value scoreboard_state attribute.
type AttributeScoreboardState int

const (
	_ AttributeScoreboardState = iota
	AttributeScoreboardStateOpen
	AttributeScoreboardStateWaiting
	AttributeScoreboardStateStarting
	AttributeScoreboardStateReading
	AttributeScoreboardStateSending
	AttributeScoreboardStateKeepalive
	AttributeScoreboardStateDnslookup
	AttributeScoreboardStateClosing
	AttributeScoreboardStateLogging
	AttributeScoreboardStateFinishing
	AttributeScoreboardStateIdleCleanup
	AttributeScoreboardStateUnknown
)

// String returns the string representation of the AttributeScoreboardState.
func (av AttributeScoreboardState) String() string {
	switch av {
	case AttributeScoreboardStateOpen:
		return "open"
	case AttributeScoreboardStateWaiting:
		return "waiting"
	case AttributeScoreboardStateStarting:
		return "starting"
	case AttributeScoreboardStateReading:
		return "reading"
	case AttributeScoreboardStateSending:
		return "sending"
	case AttributeScoreboardStateKeepalive:
		return "keepalive"
	case AttributeScoreboardStateDnslookup:
		return "dnslookup"
	case AttributeScoreboardStateClosing:
		return "closing"
	case AttributeScoreboardStateLogging:
		return "logging"
	case AttributeScoreboardStateFinishing:
		return "finishing"
	case AttributeScoreboardStateIdleCleanup:
		return "idle_cleanup"
	case AttributeScoreboardStateUnknown:
		return "unknown"
	}
	return ""
}

// MapAttributeScoreboardState is a helper map of string to AttributeScoreboardState attribute value.
var MapAttributeScoreboardState = map[string]AttributeScoreboardState{
	"open":         AttributeScoreboardStateOpen,
	"waiting":      AttributeScoreboardStateWaiting,
	"starting":     AttributeScoreboardStateStarting,
	"reading":      AttributeScoreboardStateReading,
	"sending":      AttributeScoreboardStateSending,
	"keepalive":    AttributeScoreboardStateKeepalive,
	"dnslookup":    AttributeScoreboardStateDnslookup,
	"closing":      AttributeScoreboardStateClosing,
	"logging":      AttributeScoreboardStateLogging,
	"finishing":    AttributeScoreboardStateFinishing,
	"idle_cleanup": AttributeScoreboardStateIdleCleanup,
	"unknown":      AttributeScoreboardStateUnknown,
}

// AttributeWorkersState specifies the a value workers_state attribute.
type AttributeWorkersState int

const (
	_ AttributeWorkersState = iota
	AttributeWorkersStateBusy
	AttributeWorkersStateIdle
)

// String returns the string representation of the AttributeWorkersState.
func (av AttributeWorkersState) String() string {
	switch av {
	case AttributeWorkersStateBusy:
		return "busy"
	case AttributeWorkersStateIdle:
		return "idle"
	}
	return ""
}

// MapAttributeWorkersState is a helper map of string to AttributeWorkersState attribute value.
var MapAttributeWorkersState = map[string]AttributeWorkersState{
	"busy": AttributeWorkersStateBusy,
	"idle": AttributeWorkersStateIdle,
}

type metricApacheCurrentConnections struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.current_connections metric with initial data.
func (m *metricApacheCurrentConnections) init() {
	m.data.SetName("apache.current_connections")
	m.data.SetDescription("The number of active connections currently attached to the HTTP server.")
	m.data.SetUnit("{connections}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricApacheCurrentConnections) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().InsertString("server_name", serverNameAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheCurrentConnections) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheCurrentConnections) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheCurrentConnections(settings MetricSettings) metricApacheCurrentConnections {
	m := metricApacheCurrentConnections{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheRequests struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.requests metric with initial data.
func (m *metricApacheRequests) init() {
	m.data.SetName("apache.requests")
	m.data.SetDescription("The number of requests serviced by the HTTP server per second.")
	m.data.SetUnit("{requests}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricApacheRequests) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().InsertString("server_name", serverNameAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheRequests) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheRequests) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheRequests(settings MetricSettings) metricApacheRequests {
	m := metricApacheRequests{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheScoreboard struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.scoreboard metric with initial data.
func (m *metricApacheScoreboard) init() {
	m.data.SetName("apache.scoreboard")
	m.data.SetDescription("The number of workers in each state.")
	m.data.SetUnit("{workers}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricApacheScoreboard) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string, scoreboardStateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().InsertString("server_name", serverNameAttributeValue)
	dp.Attributes().InsertString("state", scoreboardStateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheScoreboard) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheScoreboard) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheScoreboard(settings MetricSettings) metricApacheScoreboard {
	m := metricApacheScoreboard{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheTraffic struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.traffic metric with initial data.
func (m *metricApacheTraffic) init() {
	m.data.SetName("apache.traffic")
	m.data.SetDescription("Total HTTP server traffic.")
	m.data.SetUnit("By")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricApacheTraffic) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().InsertString("server_name", serverNameAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheTraffic) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheTraffic) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheTraffic(settings MetricSettings) metricApacheTraffic {
	m := metricApacheTraffic{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheUptime struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.uptime metric with initial data.
func (m *metricApacheUptime) init() {
	m.data.SetName("apache.uptime")
	m.data.SetDescription("The amount of time that the server has been running in seconds.")
	m.data.SetUnit("s")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricApacheUptime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().InsertString("server_name", serverNameAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheUptime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheUptime) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheUptime(settings MetricSettings) metricApacheUptime {
	m := metricApacheUptime{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricApacheWorkers struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills apache.workers metric with initial data.
func (m *metricApacheWorkers) init() {
	m.data.SetName("apache.workers")
	m.data.SetDescription("The number of workers currently attached to the HTTP server.")
	m.data.SetUnit("{workers}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricApacheWorkers) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, serverNameAttributeValue string, workersStateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().InsertString("server_name", serverNameAttributeValue)
	dp.Attributes().InsertString("state", workersStateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricApacheWorkers) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricApacheWorkers) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricApacheWorkers(settings MetricSettings) metricApacheWorkers {
	m := metricApacheWorkers{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                      pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                int                 // maximum observed number of metrics per resource.
	resourceCapacity               int                 // maximum observed number of resource attributes.
	metricsBuffer                  pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                      component.BuildInfo // contains version information
	metricApacheCurrentConnections metricApacheCurrentConnections
	metricApacheRequests           metricApacheRequests
	metricApacheScoreboard         metricApacheScoreboard
	metricApacheTraffic            metricApacheTraffic
	metricApacheUptime             metricApacheUptime
	metricApacheWorkers            metricApacheWorkers
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, buildInfo component.BuildInfo, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                      pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                  pmetric.NewMetrics(),
		buildInfo:                      buildInfo,
		metricApacheCurrentConnections: newMetricApacheCurrentConnections(settings.ApacheCurrentConnections),
		metricApacheRequests:           newMetricApacheRequests(settings.ApacheRequests),
		metricApacheScoreboard:         newMetricApacheScoreboard(settings.ApacheScoreboard),
		metricApacheTraffic:            newMetricApacheTraffic(settings.ApacheTraffic),
		metricApacheUptime:             newMetricApacheUptime(settings.ApacheUptime),
		metricApacheWorkers:            newMetricApacheWorkers(settings.ApacheWorkers),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).DataType() {
			case pmetric.MetricDataTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricDataTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/apachereceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricApacheCurrentConnections.emit(ils.Metrics())
	mb.metricApacheRequests.emit(ils.Metrics())
	mb.metricApacheScoreboard.emit(ils.Metrics())
	mb.metricApacheTraffic.emit(ils.Metrics())
	mb.metricApacheUptime.emit(ils.Metrics())
	mb.metricApacheWorkers.emit(ils.Metrics())
	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordApacheCurrentConnectionsDataPoint adds a data point to apache.current_connections metric.
func (mb *MetricsBuilder) RecordApacheCurrentConnectionsDataPoint(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheCurrentConnections, value was %s: %w", inputVal, err)
	}
	mb.metricApacheCurrentConnections.recordDataPoint(mb.startTime, ts, val, serverNameAttributeValue)
	return nil
}

// RecordApacheRequestsDataPoint adds a data point to apache.requests metric.
func (mb *MetricsBuilder) RecordApacheRequestsDataPoint(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheRequests, value was %s: %w", inputVal, err)
	}
	mb.metricApacheRequests.recordDataPoint(mb.startTime, ts, val, serverNameAttributeValue)
	return nil
}

// RecordApacheScoreboardDataPoint adds a data point to apache.scoreboard metric.
func (mb *MetricsBuilder) RecordApacheScoreboardDataPoint(ts pcommon.Timestamp, val int64, serverNameAttributeValue string, scoreboardStateAttributeValue AttributeScoreboardState) {
	mb.metricApacheScoreboard.recordDataPoint(mb.startTime, ts, val, serverNameAttributeValue, scoreboardStateAttributeValue.String())
}

// RecordApacheTrafficDataPoint adds a data point to apache.traffic metric.
func (mb *MetricsBuilder) RecordApacheTrafficDataPoint(ts pcommon.Timestamp, val int64, serverNameAttributeValue string) {
	mb.metricApacheTraffic.recordDataPoint(mb.startTime, ts, val, serverNameAttributeValue)
}

// RecordApacheUptimeDataPoint adds a data point to apache.uptime metric.
func (mb *MetricsBuilder) RecordApacheUptimeDataPoint(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheUptime, value was %s: %w", inputVal, err)
	}
	mb.metricApacheUptime.recordDataPoint(mb.startTime, ts, val, serverNameAttributeValue)
	return nil
}

// RecordApacheWorkersDataPoint adds a data point to apache.workers metric.
func (mb *MetricsBuilder) RecordApacheWorkersDataPoint(ts pcommon.Timestamp, inputVal string, serverNameAttributeValue string, workersStateAttributeValue AttributeWorkersState) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for ApacheWorkers, value was %s: %w", inputVal, err)
	}
	mb.metricApacheWorkers.recordDataPoint(mb.startTime, ts, val, serverNameAttributeValue, workersStateAttributeValue.String())
	return nil
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
