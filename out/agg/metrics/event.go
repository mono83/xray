package metrics

import "github.com/mono83/xray"

type definition struct {
	key  string
	args []xray.Arg
}

type event struct {
	definition
	t     xray.MetricType
	value int64
}

func (event) GetRayID() string           { return "" }
func (e event) Size() int                { return len(e.args) }
func (e event) Args() []xray.Arg         { return e.args }
func (e event) GetType() xray.MetricType { return e.t }
func (e event) GetKey() string           { return e.key }
func (e event) GetValue() int64          { return e.value }
func (e event) Get(name string) xray.Arg {
	for _, a := range e.args {
		if name == a.Name() {
			return a
		}
	}
	return nil
}
