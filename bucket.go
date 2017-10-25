package xray

// Bucket is container for arguments
type Bucket interface {
	// Size return number of args within bucket
	Size() int
	// Get returns argument (or nil) by key
	Get(string) Arg
	// Args returns args slice
	Args() []Arg
}

// CreateBucket creates bucket for provided arguments
func CreateBucket(args ... Arg) Bucket {
	l := len(args)
	if l == 0 {
		return emptyBucketInstance
	} else if l == 1 {
		return singleArgBucket{Arg: args[0]}
	}

	m := map[string]Arg{}
	for _, a := range args {
		if a != nil {
			m[a.Name()] = a
		}
	}
	return mapBucket(m)
}

// AppendBucket appends arguments to existing bucket
func AppendBucket(src Bucket, args ... Arg) Bucket {
	if src == nil && len(args) == 0 {
		return emptyBucketInstance
	} else if len(args) == 0 {
		return src
	} else if src.Size() == 0 {
		return CreateBucket(args...)
	}

	m := map[string]Arg{}
	for _, a := range src.Args() {
		if a != nil {
			m[a.Name()] = a
		}
	}
	for _, a := range args {
		if a != nil {
			m[a.Name()] = a
		}
	}
	return mapBucket(m)
}

var emptyBucketInstance = emptyBucket{}

type emptyBucket struct {
}

func (emptyBucket) Size() int      { return 0 }
func (emptyBucket) Get(string) Arg { return nil }
func (emptyBucket) Args() []Arg    { return nil }

type singleArgBucket struct {
	Arg
}

func (s singleArgBucket) Size() int { return 1 }
func (s singleArgBucket) Get(key string) Arg {
	if s.Arg.Name() == key {
		return s.Arg
	}
	return nil
}
func (s singleArgBucket) Args() []Arg { return []Arg{s.Arg} }

type mapBucket map[string]Arg

func (m mapBucket) Size() int { return len(m) }
func (m mapBucket) Get(key string) Arg {
	if v, ok := m[key]; ok {
		return v
	}

	return nil
}
func (m mapBucket) Args() []Arg {
	args := []Arg{}
	for _, v := range m {
		args = append(args, v)
	}
	return args
}
