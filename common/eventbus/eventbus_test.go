package eventbus

import "testing"

func TestEventBus(t *testing.T) {
	eb := NewEventBus()
	received := false

	// Subscribe to an event
	eb.Subscribe("test_event", func(data interface{}) interface{} {
		received = true
		return nil
	})

	// Publish the event
	eb.Publish("test_event", nil)

	if !received {
		t.Error("Expected event handler to be called")
	}
}

func TestEventBusReturnValue(t *testing.T) {
	eb := NewEventBus()
	var returnedValue interface{}

	// Subscribe to an event
	eb.Subscribe("return_event", func(data interface{}) interface{} {
		return "Returned Value"
	})

	// Publish the event
	returnedValue = eb.Publish("return_event", nil)

	if returnedValue != "Returned Value" {
		t.Errorf("Expected 'Returned Value', got %v", returnedValue)
	}
}
