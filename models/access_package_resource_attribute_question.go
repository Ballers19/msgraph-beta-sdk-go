package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceAttributeQuestion 
type AccessPackageResourceAttributeQuestion struct {
    AccessPackageResourceAttributeSource
    // The question asked in order to get the value of the attribute
    question AccessPackageQuestionable
}
// NewAccessPackageResourceAttributeQuestion instantiates a new AccessPackageResourceAttributeQuestion and sets the default values.
func NewAccessPackageResourceAttributeQuestion()(*AccessPackageResourceAttributeQuestion) {
    m := &AccessPackageResourceAttributeQuestion{
        AccessPackageResourceAttributeSource: *NewAccessPackageResourceAttributeSource(),
    }
    return m
}
// CreateAccessPackageResourceAttributeQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageResourceAttributeQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceAttributeQuestion(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceAttributeQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessPackageResourceAttributeSource.GetFieldDeserializers()
    res["question"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuestion(val.(AccessPackageQuestionable))
        }
        return nil
    }
    return res
}
// GetQuestion gets the question property value. The question asked in order to get the value of the attribute
func (m *AccessPackageResourceAttributeQuestion) GetQuestion()(AccessPackageQuestionable) {
    if m == nil {
        return nil
    } else {
        return m.question
    }
}
// Serialize serializes information the current object
func (m *AccessPackageResourceAttributeQuestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessPackageResourceAttributeSource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("question", m.GetQuestion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetQuestion sets the question property value. The question asked in order to get the value of the attribute
func (m *AccessPackageResourceAttributeQuestion) SetQuestion(value AccessPackageQuestionable)() {
    if m != nil {
        m.question = value
    }
}
