package v2

import (
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
	v1 "github.com/maysunfaisal/api/v2/pkg/apis/workspaces/v1alpha2"
)

// GetEvents returns the Events Object parsed from devfile
func (d *DevfileV2) GetEvents() v1.Events {
	if d.Events != nil {
		return *d.Events
	}
	return v1.Events{}
}

// AddEvents adds the Events Object to the devfile's events
// if the event is already defined in the devfile, error out
func (d *DevfileV2) AddEvents(events v1.Events) error {

	if d.Events == nil {
		d.Events = &v1.Events{}
	}

	if len(events.PreStop) > 0 {
		if len(d.Events.PreStop) > 0 {
			return &common.FieldAlreadyExistError{Field: "pre stop"}
		}
		d.Events.PreStop = events.PreStop
	}

	if len(events.PreStart) > 0 {
		if len(d.Events.PreStart) > 0 {
			return &common.FieldAlreadyExistError{Field: "pre start"}
		}
		d.Events.PreStart = events.PreStart
	}

	if len(events.PostStop) > 0 {
		if len(d.Events.PostStop) > 0 {
			return &common.FieldAlreadyExistError{Field: "post stop"}
		}
		d.Events.PostStop = events.PostStop
	}

	if len(events.PostStart) > 0 {
		if len(d.Events.PostStart) > 0 {
			return &common.FieldAlreadyExistError{Field: "post start"}
		}
		d.Events.PostStart = events.PostStart
	}

	return nil
}

// UpdateEvents updates the devfile's events
// it only updates the events passed to it
func (d *DevfileV2) UpdateEvents(postStart, postStop, preStart, preStop []string) {

	if d.Events == nil {
		d.Events = &v1.Events{}
	}

	if postStart != nil {
		d.Events.PostStart = postStart
	}
	if postStop != nil {
		d.Events.PostStop = postStop
	}
	if preStart != nil {
		d.Events.PreStart = preStart
	}
	if preStop != nil {
		d.Events.PreStop = preStop
	}
}
