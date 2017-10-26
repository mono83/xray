package xray

// newMetricEvent builds new metric event
func newMetricEvent(r Ray, t MetricType, key string, value int64, args []Arg) *metricEvent {
	return &metricEvent{
		event: event{
			id:       r.GetRayID(),
			base:     r.GetArguments(),
			provided: args,
		},
		t:     t,
		value: value,
		key:   r.GetMetricPrefix() + key,
	}
}

// newLogEvent builds new logging event
func newLogEvent(r Ray, t Level, message string, args []Arg) LogEvent {
	return &logEvent{
		event: event{
			id:       r.GetRayID(),
			base:     r.GetArguments(),
			provided: args,
		},
		t:       t,
		logger:  r.GetLogger(),
		message: message,
	}
}

type event struct {
	id         string
	base       Bucket
	calculated Bucket
	provided   []Arg
}

func (e event) GetRayID() string    { return e.id }
func (e *event) Size() int          { return e.Bucket().Size() }
func (e *event) Get(key string) Arg { return e.Bucket().Get(key) }
func (e *event) Args() []Arg        { return e.Bucket().Args() }

func (e *event) Bucket() Bucket {
	if e.calculated == nil {
		e.calculated = AppendBucket(e.base, e.provided...)
	}

	return e.calculated
}

type metricEvent struct {
	event
	t     MetricType
	key   string
	value int64
}

func (m metricEvent) GetType() MetricType { return m.t }
func (m metricEvent) GetKey() string      { return m.key }
func (m metricEvent) GetValue() int64     { return m.value }

type logEvent struct {
	event
	t       Level
	logger  string
	message string
}

func (l logEvent) GetLevel() Level    { return l.t }
func (l logEvent) GetLogger() string  { return l.logger }
func (l logEvent) GetMessage() string { return l.message }
