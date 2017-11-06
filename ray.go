package xray

import "github.com/mono83/xray/args"

// New creates new root-level ray
func New(emitterGenerator func() EventEmitter, rayIDGenerator func() string) Ray {
	rayID := rayIDGenerator()
	return &ray{
		id: rayID,

		bucket:       singleArgBucket{args.RayID(rayID)},
		EventEmitter: emitterGenerator(),

		genRayID:   rayIDGenerator,
		getEmitter: emitterGenerator,
	}
}

type ray struct {
	EventEmitter
	id           string
	logger       string
	metricPrefix string
	bucket       Bucket
	genRayID     func() string
	getEmitter   func() EventEmitter
}

func (r *ray) GetRayID() string        { return r.id }
func (r *ray) GetArguments() Bucket    { return r.bucket }
func (r *ray) GetLogger() string       { return r.logger }
func (r *ray) GetMetricPrefix() string { return r.metricPrefix }
func (r *ray) WithLogger(v string) Ray {
	c := r.clone()
	c.logger = v
	return c
}
func (r *ray) WithMetricPrefix(v string) Ray {
	if v == "" {
		return r
	}
	if v[len(v) - 1] != '.' {
		v += "."
	}
	c := r.clone()
	c.metricPrefix = c.metricPrefix + v
	return c
}
func (r *ray) With(args ...Arg) Ray {
	if len(args) == 0 {
		return r
	}

	c := r.clone()
	c.bucket = AppendBucket(c.bucket, args...)
	return c
}
func (r *ray) Fork() Ray {
	c := r.clone()
	c.id = c.genRayID()
	c.EventEmitter = c.getEmitter()
	c.EventEmitter.On(func(ee ...Event) {
		r.EventEmitter.Emit(ee...)
	})
	if c.bucket.Size() < 2 {
		c.bucket = singleArgBucket{args.RayID(c.id)}
	} else {
		c.bucket = AppendBucket(c.bucket, args.RayID(c.id))
	}
	return c
}
func (r *ray) clone() *ray {
	return &ray{
		EventEmitter: r.EventEmitter,
		id:           r.id,
		bucket:       r.bucket,
		logger:       r.logger,
		metricPrefix: r.metricPrefix,

		genRayID:   r.genRayID,
		getEmitter: r.getEmitter,
	}
}
