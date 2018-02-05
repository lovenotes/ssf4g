package tlog

type LogEvent struct {
	_event_id string
	_ret_ch   chan error
}

func NewLogEvent(eventid string, retch chan error) *LogEvent {
	return &LogEvent{
		_event_id: eventid,
		_ret_ch:   retch,
	}
}

func (logevent *LogEvent) Wait() {
	select {
	case <-logevent._ret_ch:
	}
}
