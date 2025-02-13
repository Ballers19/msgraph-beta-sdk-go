package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// RetentionEvent provides operations to manage the collection of activityStatistics entities.
type RetentionEvent struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The user who created the retentionEvent.
    createdBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The date time when the retentionEvent was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Optional information about the event.
    description *string
    // Name of the event.
    displayName *string
    // The eventPropagationResults property
    eventPropagationResults []EventPropagationResultable
    // Represents the workload (SharePoint Online, OneDrive for Business, Exchange Online) and identification information associated with a retention event.
    eventQueries []EventQueryable
    // The eventStatus property
    eventStatus RetentionEventStatusable
    // Optional time when the event should be triggered.
    eventTriggerDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user who last modified the retentionEvent.
    lastModifiedBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The latest date time when the retentionEvent was modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last time the status of the event was updated.
    lastStatusUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies the event that will start the retention period for labels that use this event type when an event is created.
    retentionEventType RetentionEventTypeable
}
// NewRetentionEvent instantiates a new retentionEvent and sets the default values.
func NewRetentionEvent()(*RetentionEvent) {
    m := &RetentionEvent{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateRetentionEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetentionEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionEvent(), nil
}
// GetCreatedBy gets the createdBy property value. The user who created the retentionEvent.
func (m *RetentionEvent) GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date time when the retentionEvent was created.
func (m *RetentionEvent) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Optional information about the event.
func (m *RetentionEvent) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Name of the event.
func (m *RetentionEvent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEventPropagationResults gets the eventPropagationResults property value. The eventPropagationResults property
func (m *RetentionEvent) GetEventPropagationResults()([]EventPropagationResultable) {
    if m == nil {
        return nil
    } else {
        return m.eventPropagationResults
    }
}
// GetEventQueries gets the eventQueries property value. Represents the workload (SharePoint Online, OneDrive for Business, Exchange Online) and identification information associated with a retention event.
func (m *RetentionEvent) GetEventQueries()([]EventQueryable) {
    if m == nil {
        return nil
    } else {
        return m.eventQueries
    }
}
// GetEventStatus gets the eventStatus property value. The eventStatus property
func (m *RetentionEvent) GetEventStatus()(RetentionEventStatusable) {
    if m == nil {
        return nil
    } else {
        return m.eventStatus
    }
}
// GetEventTriggerDateTime gets the eventTriggerDateTime property value. Optional time when the event should be triggered.
func (m *RetentionEvent) GetEventTriggerDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventTriggerDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetentionEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["eventPropagationResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventPropagationResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EventPropagationResultable, len(val))
            for i, v := range val {
                res[i] = v.(EventPropagationResultable)
            }
            m.SetEventPropagationResults(res)
        }
        return nil
    }
    res["eventQueries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventQueryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EventQueryable, len(val))
            for i, v := range val {
                res[i] = v.(EventQueryable)
            }
            m.SetEventQueries(res)
        }
        return nil
    }
    res["eventStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRetentionEventStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventStatus(val.(RetentionEventStatusable))
        }
        return nil
    }
    res["eventTriggerDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventTriggerDateTime(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["lastStatusUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastStatusUpdateDateTime(val)
        }
        return nil
    }
    res["retentionEventType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRetentionEventTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetentionEventType(val.(RetentionEventTypeable))
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the retentionEvent.
func (m *RetentionEvent) GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The latest date time when the retentionEvent was modified.
func (m *RetentionEvent) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLastStatusUpdateDateTime gets the lastStatusUpdateDateTime property value. Last time the status of the event was updated.
func (m *RetentionEvent) GetLastStatusUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastStatusUpdateDateTime
    }
}
// GetRetentionEventType gets the retentionEventType property value. Specifies the event that will start the retention period for labels that use this event type when an event is created.
func (m *RetentionEvent) GetRetentionEventType()(RetentionEventTypeable) {
    if m == nil {
        return nil
    } else {
        return m.retentionEventType
    }
}
// Serialize serializes information the current object
func (m *RetentionEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEventPropagationResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEventPropagationResults()))
        for i, v := range m.GetEventPropagationResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("eventPropagationResults", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEventQueries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEventQueries()))
        for i, v := range m.GetEventQueries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("eventQueries", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("eventStatus", m.GetEventStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventTriggerDateTime", m.GetEventTriggerDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastStatusUpdateDateTime", m.GetLastStatusUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("retentionEventType", m.GetRetentionEventType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The user who created the retentionEvent.
func (m *RetentionEvent) SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date time when the retentionEvent was created.
func (m *RetentionEvent) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Optional information about the event.
func (m *RetentionEvent) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Name of the event.
func (m *RetentionEvent) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEventPropagationResults sets the eventPropagationResults property value. The eventPropagationResults property
func (m *RetentionEvent) SetEventPropagationResults(value []EventPropagationResultable)() {
    if m != nil {
        m.eventPropagationResults = value
    }
}
// SetEventQueries sets the eventQueries property value. Represents the workload (SharePoint Online, OneDrive for Business, Exchange Online) and identification information associated with a retention event.
func (m *RetentionEvent) SetEventQueries(value []EventQueryable)() {
    if m != nil {
        m.eventQueries = value
    }
}
// SetEventStatus sets the eventStatus property value. The eventStatus property
func (m *RetentionEvent) SetEventStatus(value RetentionEventStatusable)() {
    if m != nil {
        m.eventStatus = value
    }
}
// SetEventTriggerDateTime sets the eventTriggerDateTime property value. Optional time when the event should be triggered.
func (m *RetentionEvent) SetEventTriggerDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.eventTriggerDateTime = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the retentionEvent.
func (m *RetentionEvent) SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The latest date time when the retentionEvent was modified.
func (m *RetentionEvent) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLastStatusUpdateDateTime sets the lastStatusUpdateDateTime property value. Last time the status of the event was updated.
func (m *RetentionEvent) SetLastStatusUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastStatusUpdateDateTime = value
    }
}
// SetRetentionEventType sets the retentionEventType property value. Specifies the event that will start the retention period for labels that use this event type when an event is created.
func (m *RetentionEvent) SetRetentionEventType(value RetentionEventTypeable)() {
    if m != nil {
        m.retentionEventType = value
    }
}
