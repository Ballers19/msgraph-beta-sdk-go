package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BufferDecryptionResult 
type BufferDecryptionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The decryptedBuffer property
    decryptedBuffer []byte
}
// NewBufferDecryptionResult instantiates a new bufferDecryptionResult and sets the default values.
func NewBufferDecryptionResult()(*BufferDecryptionResult) {
    m := &BufferDecryptionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBufferDecryptionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBufferDecryptionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBufferDecryptionResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BufferDecryptionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDecryptedBuffer gets the decryptedBuffer property value. The decryptedBuffer property
func (m *BufferDecryptionResult) GetDecryptedBuffer()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.decryptedBuffer
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BufferDecryptionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["decryptedBuffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecryptedBuffer(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *BufferDecryptionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("decryptedBuffer", m.GetDecryptedBuffer())
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
func (m *BufferDecryptionResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDecryptedBuffer sets the decryptedBuffer property value. The decryptedBuffer property
func (m *BufferDecryptionResult) SetDecryptedBuffer(value []byte)() {
    if m != nil {
        m.decryptedBuffer = value
    }
}
