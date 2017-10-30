package out

import (
	"github.com/mono83/xray"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitter_Empty(t *testing.T) {
	assert := assert2.New(t)

	cntInvoke := 0
	cntEvents := 0
	spl := Splitter(func(ee ...xray.Event) {
		cntInvoke++
		cntEvents += len(ee)
	}, 2)

	spl()
	assert.Equal(0, cntInvoke)
	assert.Equal(0, cntEvents)

	spl([]xray.Event{}...)
	assert.Equal(0, cntInvoke)
	assert.Equal(0, cntEvents)
}

func TestSplitter_Below(t *testing.T) {
	assert := assert2.New(t)

	cntInvoke := 0
	cntEvents := 0
	spl := Splitter(func(ee ...xray.Event) {
		cntInvoke++
		cntEvents += len(ee)
	}, 3)

	spl(event{})
	assert.Equal(1, cntInvoke)
	assert.Equal(1, cntEvents)

	cntInvoke = 0
	cntEvents = 0

	spl(event{}, event{})
	assert.Equal(1, cntInvoke)
	assert.Equal(2, cntEvents)
}

func TestSplitter_Exact(t *testing.T) {
	assert := assert2.New(t)

	cntInvoke := 0
	cntEvents := 0
	spl := Splitter(func(ee ...xray.Event) {
		cntInvoke++
		cntEvents += len(ee)
	}, 3)

	spl(event{}, event{}, event{})
	assert.Equal(1, cntInvoke)
	assert.Equal(3, cntEvents)

	cntInvoke = 0
	cntEvents = 0

	spl(event{}, event{}, event{}, event{}, event{}, event{})
	assert.Equal(2, cntInvoke)
	assert.Equal(6, cntEvents)
}

func TestSplitter_Above(t *testing.T) {
	assert := assert2.New(t)

	cntInvoke := 0
	cntEvents := 0
	spl := Splitter(func(ee ...xray.Event) {
		cntInvoke++
		cntEvents += len(ee)
	}, 2)

	spl(event{}, event{}, event{})
	assert.Equal(2, cntInvoke)
	assert.Equal(3, cntEvents)

	cntInvoke = 0
	cntEvents = 0

	spl(event{}, event{}, event{}, event{}, event{})
	assert.Equal(3, cntInvoke)
	assert.Equal(5, cntEvents)
}

type event struct{}

func (event) Args() []xray.Arg    { return nil }
func (event) Get(string) xray.Arg { return nil }
func (event) Size() int           { return 0 }
func (event) GetRayID() string    { return "" }
