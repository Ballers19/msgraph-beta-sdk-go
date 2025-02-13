package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsOverview 
type UserExperienceAnalyticsOverview struct {
    Entity
    // The user experience analytics insights.
    insights []UserExperienceAnalyticsInsightable
}
// NewUserExperienceAnalyticsOverview instantiates a new UserExperienceAnalyticsOverview and sets the default values.
func NewUserExperienceAnalyticsOverview()(*UserExperienceAnalyticsOverview) {
    m := &UserExperienceAnalyticsOverview{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsOverview(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["insights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsInsightFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsInsightable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsInsightable)
            }
            m.SetInsights(res)
        }
        return nil
    }
    return res
}
// GetInsights gets the insights property value. The user experience analytics insights.
func (m *UserExperienceAnalyticsOverview) GetInsights()([]UserExperienceAnalyticsInsightable) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetInsights() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInsights()))
        for i, v := range m.GetInsights() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("insights", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInsights sets the insights property value. The user experience analytics insights.
func (m *UserExperienceAnalyticsOverview) SetInsights(value []UserExperienceAnalyticsInsightable)() {
    if m != nil {
        m.insights = value
    }
}
