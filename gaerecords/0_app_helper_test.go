package gaerecords

import (
	"os"
	"testing"
	"appengine/datastore"
	"gae-go-testing.googlecode.com/git/appenginetesting"
)

// Creates a test appengine context object
func UseTestAppEngineContext() {
	
	appEngineContext, _ = appenginetesting.NewContext(nil)
	
}

func CreateTestRecord() *Record {
	return NewRecord(CreateTestModel())
}

func CreateTestModel() *Model {
	return NewModel("model")
}

func CreatePersistedRecord(model *Model) (*Record, os.Error) {
	
	context := GetAppEngineContext()
	people := CreateTestModel()
	person := people.New()
	key := person.DatastoreKey()
	
	person.Set("name", "Mat").Set("age", int64(29))
	
	newKey, err := datastore.Put(context, key, datastore.PropertyLoadSaver(person))
	
	if err != nil {
		return nil, err
	}
	
	// set the person ID
	person.setID(newKey.IntID())
	
	return person, nil
	
}


func TestSetup(t *testing.T) {
	
	UseTestAppEngineContext()
	t.Logf("<<< Test context created %v >>>", appEngineContext)
	
}