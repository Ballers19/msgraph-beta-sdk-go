package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ToneInfo 
type ToneInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // An incremental identifier used for ordering DTMF events.
    sequenceId *int64
    // The tone property
    tone *Tone
}
// NewToneInfo instantiates a new toneInfo and sets the default values.
func NewToneInfo()(*ToneInfo) {
    m := &ToneInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateToneInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateToneInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewToneInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ToneInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ToneInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sequenceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequenceId(val)
        }
        return nil
    }
    res["tone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTone)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTone(val.(*Tone))
        }
        return nil
    }
    return res
}
// GetSequenceId gets the sequenceId property value. An incremental identifier used for ordering DTMF events.
func (m *ToneInfo) GetSequenceId()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sequenceId
    }
}
// GetTone gets the tone property value. The tone property
func (m *ToneInfo) GetTone()(*Tone) {
    if m == nil {
        return nil
    } else {
        return m.tone
    }
}
// Serialize serializes information the current object
func (m *ToneInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("sequenceId", m.GetSequenceId())
        if err != nil {
            return err
        }
    }
    if m.GetTone() != nil {
        cast := (*m.GetTone()).String()
        err := writer.WriteStringValue("tone", &cast)
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
func (m *ToneInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSequenceId sets the sequenceId property value. An incremental identifier used for ordering DTMF events.
func (m *ToneInfo) SetSequenceId(value *int64)() {
    if m != nil {
        m.sequenceId = value
    }
}
// SetTone sets the tone property value. The tone property
func (m *ToneInfo) SetTone(value *Tone)() {
    if m != nil {
        m.tone = value
    }
}
