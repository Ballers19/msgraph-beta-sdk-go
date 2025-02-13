package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VerificationResult 
type VerificationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The signatureValid property
    signatureValid *bool
}
// NewVerificationResult instantiates a new verificationResult and sets the default values.
func NewVerificationResult()(*VerificationResult) {
    m := &VerificationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVerificationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVerificationResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerificationResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerificationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerificationResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["signatureValid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignatureValid(val)
        }
        return nil
    }
    return res
}
// GetSignatureValid gets the signatureValid property value. The signatureValid property
func (m *VerificationResult) GetSignatureValid()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.signatureValid
    }
}
// Serialize serializes information the current object
func (m *VerificationResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("signatureValid", m.GetSignatureValid())
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
func (m *VerificationResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSignatureValid sets the signatureValid property value. The signatureValid property
func (m *VerificationResult) SetSignatureValid(value *bool)() {
    if m != nil {
        m.signatureValid = value
    }
}
