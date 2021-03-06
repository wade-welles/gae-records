package gaerecords

import (
	"testing"
)

func TestNewModel(t *testing.T) {

	model := NewModel("kind1")

	assertEqual(t, model.RecordType(), "kind1")

}

func TestNewModelWithFunc(t *testing.T) {

	var called bool = false
	var lastModel *Model = nil

	model := NewModel("kind2", func(m *Model) {

		called = true
		lastModel = m

	})

	assertEqual(t, true, called)
	assertEqual(t, model, lastModel)

}

func TestNew(t *testing.T) {

	model := CreateTestModel()
	record := model.New()

	assertNotNil(t, record, "Record shouldn't be nil")

}

func TestModelString(t *testing.T) {

	model := NewModel("people")

	assertEqual(t, "{Model:people}", model.String())

}
