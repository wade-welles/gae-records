package gaerecords

import (
	"testing"
)

func TestModelAfterFindEvent(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("afterFindEventModel")
	record, _ := CreatePersistedRecord(t, model)
	
	var called bool = false
	var context *EventContext = nil
	
	model.AfterFind.Do(func(c *EventContext){
		called = true
		context = c
	})
	
	// do something that should trigger the event
	model.Find(record.ID())
	
	assertEqual(t, true, called)
	assertNotNil(t, context.Args[0], "context.Args[0]")
	assertEqual(t, record.ID(), context.Args[0].(*Record).ID())
	
}

func TestModelAfterFindEvent_withFindAll(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("afterFindEventALLModel")
	CreatePersistedRecord(t, model)
	record2, _ := CreatePersistedRecord(t, model)
	
	var called int = 0
	var lastContext *EventContext = nil
	
	model.AfterFind.Do(func(c *EventContext){
		called++
		lastContext = c
	})
	
	// do something that should trigger the event
	model.All()
	
	assertEqual(t, 2, called)
	assertNotNil(t, lastContext.Args[0], "context.Args[0]")
	assertEqual(t, record2.ID(), lastContext.Args[0].(*Record).ID())
	
}

func TestModelBeforeDeleteByIDEvent(t *testing.T) {

	model := CreateTestModelWithPropertyType("afterFindEventModel")
	record, _ := CreatePersistedRecord(t, model)
	
	var called bool = false
	var context *EventContext = nil
	
	model.BeforeDeleteByID.Do(func(c *EventContext){
		called = true
		context = c
	})
	
	// do something that should trigger the event
	model.Delete(record.ID())
	
	assertEqual(t, true, called)
	assertNotNil(t, context.Args[0], "context.Args[0]")
	assertEqual(t, record.ID(), context.Args[0].(int64))
	
}

func TestModelBeforeDeleteByIDEvent_Cancellation(t *testing.T) {

	model := CreateTestModelWithPropertyType("beforeDeleteEventModel")
	record, _ := CreatePersistedRecord(t, model)
	
	var called bool = false
	var context *EventContext = nil
	
	model.BeforeDeleteByID.Do(func(c *EventContext){
		called = true
		context = c
		
		// cancel the delete
		c.Cancel = true
		
	})
	
	// do something that should trigger the event
	err := model.Delete(record.ID())
	
	if err != ErrOperationCancelledByEventCallback {
		t.Errorf("ErrOperationCancelledByEventCallback Error should be returned if the delete operation was cancelled by an event callback")
	}
	
	foundRecord, _ := model.Find(record.ID())
	
	assertEqual(t, record.ID(), foundRecord.ID())
	
}