package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EvaluateLabelJobResultGroup 
type EvaluateLabelJobResultGroup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The automatic property
    automatic EvaluateLabelJobResultable
    // The recommended property
    recommended EvaluateLabelJobResultable
}
// NewEvaluateLabelJobResultGroup instantiates a new evaluateLabelJobResultGroup and sets the default values.
func NewEvaluateLabelJobResultGroup()(*EvaluateLabelJobResultGroup) {
    m := &EvaluateLabelJobResultGroup{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEvaluateLabelJobResultGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateLabelJobResultGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEvaluateLabelJobResultGroup(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateLabelJobResultGroup) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAutomatic gets the automatic property value. The automatic property
func (m *EvaluateLabelJobResultGroup) GetAutomatic()(EvaluateLabelJobResultable) {
    if m == nil {
        return nil
    } else {
        return m.automatic
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateLabelJobResultGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["automatic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEvaluateLabelJobResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomatic(val.(EvaluateLabelJobResultable))
        }
        return nil
    }
    res["recommended"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEvaluateLabelJobResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommended(val.(EvaluateLabelJobResultable))
        }
        return nil
    }
    return res
}
// GetRecommended gets the recommended property value. The recommended property
func (m *EvaluateLabelJobResultGroup) GetRecommended()(EvaluateLabelJobResultable) {
    if m == nil {
        return nil
    } else {
        return m.recommended
    }
}
// Serialize serializes information the current object
func (m *EvaluateLabelJobResultGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("automatic", m.GetAutomatic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recommended", m.GetRecommended())
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
func (m *EvaluateLabelJobResultGroup) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAutomatic sets the automatic property value. The automatic property
func (m *EvaluateLabelJobResultGroup) SetAutomatic(value EvaluateLabelJobResultable)() {
    if m != nil {
        m.automatic = value
    }
}
// SetRecommended sets the recommended property value. The recommended property
func (m *EvaluateLabelJobResultGroup) SetRecommended(value EvaluateLabelJobResultable)() {
    if m != nil {
        m.recommended = value
    }
}
