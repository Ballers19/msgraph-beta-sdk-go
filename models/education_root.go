package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationRoot 
type EducationRoot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The classes property
    classes []EducationClassable
    // The me property
    me EducationUserable
    // The schools property
    schools []EducationSchoolable
    // The synchronizationProfiles property
    synchronizationProfiles []EducationSynchronizationProfileable
    // The users property
    users []EducationUserable
}
// NewEducationRoot instantiates a new EducationRoot and sets the default values.
func NewEducationRoot()(*EducationRoot) {
    m := &EducationRoot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationRoot(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationRoot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClasses gets the classes property value. The classes property
func (m *EducationRoot) GetClasses()([]EducationClassable) {
    if m == nil {
        return nil
    } else {
        return m.classes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationClassFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationClassable, len(val))
            for i, v := range val {
                res[i] = v.(EducationClassable)
            }
            m.SetClasses(res)
        }
        return nil
    }
    res["me"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMe(val.(EducationUserable))
        }
        return nil
    }
    res["schools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSchoolFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSchoolable, len(val))
            for i, v := range val {
                res[i] = v.(EducationSchoolable)
            }
            m.SetSchools(res)
        }
        return nil
    }
    res["synchronizationProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSynchronizationProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSynchronizationProfileable, len(val))
            for i, v := range val {
                res[i] = v.(EducationSynchronizationProfileable)
            }
            m.SetSynchronizationProfiles(res)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationUserable, len(val))
            for i, v := range val {
                res[i] = v.(EducationUserable)
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetMe gets the me property value. The me property
func (m *EducationRoot) GetMe()(EducationUserable) {
    if m == nil {
        return nil
    } else {
        return m.me
    }
}
// GetSchools gets the schools property value. The schools property
func (m *EducationRoot) GetSchools()([]EducationSchoolable) {
    if m == nil {
        return nil
    } else {
        return m.schools
    }
}
// GetSynchronizationProfiles gets the synchronizationProfiles property value. The synchronizationProfiles property
func (m *EducationRoot) GetSynchronizationProfiles()([]EducationSynchronizationProfileable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationProfiles
    }
}
// GetUsers gets the users property value. The users property
func (m *EducationRoot) GetUsers()([]EducationUserable) {
    if m == nil {
        return nil
    } else {
        return m.users
    }
}
// Serialize serializes information the current object
func (m *EducationRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClasses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClasses()))
        for i, v := range m.GetClasses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("classes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("me", m.GetMe())
        if err != nil {
            return err
        }
    }
    if m.GetSchools() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSchools()))
        for i, v := range m.GetSchools() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("schools", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSynchronizationProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSynchronizationProfiles()))
        for i, v := range m.GetSynchronizationProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("synchronizationProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("users", cast)
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
func (m *EducationRoot) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClasses sets the classes property value. The classes property
func (m *EducationRoot) SetClasses(value []EducationClassable)() {
    if m != nil {
        m.classes = value
    }
}
// SetMe sets the me property value. The me property
func (m *EducationRoot) SetMe(value EducationUserable)() {
    if m != nil {
        m.me = value
    }
}
// SetSchools sets the schools property value. The schools property
func (m *EducationRoot) SetSchools(value []EducationSchoolable)() {
    if m != nil {
        m.schools = value
    }
}
// SetSynchronizationProfiles sets the synchronizationProfiles property value. The synchronizationProfiles property
func (m *EducationRoot) SetSynchronizationProfiles(value []EducationSynchronizationProfileable)() {
    if m != nil {
        m.synchronizationProfiles = value
    }
}
// SetUsers sets the users property value. The users property
func (m *EducationRoot) SetUsers(value []EducationUserable)() {
    if m != nil {
        m.users = value
    }
}
