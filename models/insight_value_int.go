package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InsightValueInt 
type InsightValueInt struct {
    UserExperienceAnalyticsInsightValue
    // Not yet documented
    value *int32
}
// NewInsightValueInt instantiates a new InsightValueInt and sets the default values.
func NewInsightValueInt()(*InsightValueInt) {
    m := &InsightValueInt{
        UserExperienceAnalyticsInsightValue: *NewUserExperienceAnalyticsInsightValue(),
    }
    return m
}
// CreateInsightValueIntFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInsightValueIntFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInsightValueInt(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InsightValueInt) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserExperienceAnalyticsInsightValue.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. Not yet documented
func (m *InsightValueInt) GetValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *InsightValueInt) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserExperienceAnalyticsInsightValue.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. Not yet documented
func (m *InsightValueInt) SetValue(value *int32)() {
    if m != nil {
        m.value = value
    }
}
