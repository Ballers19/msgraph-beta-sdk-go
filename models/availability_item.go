package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AvailabilityItem 
type AvailabilityItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The endDateTime property
    endDateTime DateTimeTimeZoneable
    // Indicates the service ID in case of 1:n appointments. If the appointment is of type 1:n, this field will be present, otherwise, null.
    serviceId *string
    // The startDateTime property
    startDateTime DateTimeTimeZoneable
    // The status of the staff member. Possible values are: available, busy, slotsAvailable, outOfOffice, unknownFutureValue.
    status *BookingsAvailabilityStatus
}
// NewAvailabilityItem instantiates a new availabilityItem and sets the default values.
func NewAvailabilityItem()(*AvailabilityItem) {
    m := &AvailabilityItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAvailabilityItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAvailabilityItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAvailabilityItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AvailabilityItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *AvailabilityItem) GetEndDateTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AvailabilityItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["serviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceId(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingsAvailabilityStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*BookingsAvailabilityStatus))
        }
        return nil
    }
    return res
}
// GetServiceId gets the serviceId property value. Indicates the service ID in case of 1:n appointments. If the appointment is of type 1:n, this field will be present, otherwise, null.
func (m *AvailabilityItem) GetServiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceId
    }
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *AvailabilityItem) GetStartDateTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetStatus gets the status property value. The status of the staff member. Possible values are: available, busy, slotsAvailable, outOfOffice, unknownFutureValue.
func (m *AvailabilityItem) GetStatus()(*BookingsAvailabilityStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Serialize serializes information the current object
func (m *AvailabilityItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serviceId", m.GetServiceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AvailabilityItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *AvailabilityItem) SetEndDateTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetServiceId sets the serviceId property value. Indicates the service ID in case of 1:n appointments. If the appointment is of type 1:n, this field will be present, otherwise, null.
func (m *AvailabilityItem) SetServiceId(value *string)() {
    if m != nil {
        m.serviceId = value
    }
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *AvailabilityItem) SetStartDateTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetStatus sets the status property value. The status of the staff member. Possible values are: available, busy, slotsAvailable, outOfOffice, unknownFutureValue.
func (m *AvailabilityItem) SetStatus(value *BookingsAvailabilityStatus)() {
    if m != nil {
        m.status = value
    }
}
