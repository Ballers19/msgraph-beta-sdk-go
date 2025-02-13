package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequestAllSiteLocation 
type SubjectRightsRequestAllSiteLocation struct {
    SubjectRightsRequestSiteLocation
}
// NewSubjectRightsRequestAllSiteLocation instantiates a new SubjectRightsRequestAllSiteLocation and sets the default values.
func NewSubjectRightsRequestAllSiteLocation()(*SubjectRightsRequestAllSiteLocation) {
    m := &SubjectRightsRequestAllSiteLocation{
        SubjectRightsRequestSiteLocation: *NewSubjectRightsRequestSiteLocation(),
    }
    return m
}
// CreateSubjectRightsRequestAllSiteLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubjectRightsRequestAllSiteLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequestAllSiteLocation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequestAllSiteLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubjectRightsRequestSiteLocation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SubjectRightsRequestAllSiteLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubjectRightsRequestSiteLocation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
