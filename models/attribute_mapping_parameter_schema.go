package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeMappingParameterSchema 
type AttributeMappingParameterSchema struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The given parameter can be provided multiple times (for example, multiple input strings in the Concatenate(string,string,...) function).
    allowMultipleOccurrences *bool
    // Parameter name.
    name *string
    // true if the parameter is required; otherwise false.
    required *bool
    // The type property
    type_escaped *AttributeType
}
// NewAttributeMappingParameterSchema instantiates a new attributeMappingParameterSchema and sets the default values.
func NewAttributeMappingParameterSchema()(*AttributeMappingParameterSchema) {
    m := &AttributeMappingParameterSchema{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAttributeMappingParameterSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeMappingParameterSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeMappingParameterSchema(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttributeMappingParameterSchema) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowMultipleOccurrences gets the allowMultipleOccurrences property value. The given parameter can be provided multiple times (for example, multiple input strings in the Concatenate(string,string,...) function).
func (m *AttributeMappingParameterSchema) GetAllowMultipleOccurrences()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleOccurrences
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMappingParameterSchema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowMultipleOccurrences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMultipleOccurrences(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*AttributeType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Parameter name.
func (m *AttributeMappingParameterSchema) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetRequired gets the required property value. true if the parameter is required; otherwise false.
func (m *AttributeMappingParameterSchema) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
// GetType gets the type property value. The type property
func (m *AttributeMappingParameterSchema) GetType()(*AttributeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *AttributeMappingParameterSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleOccurrences", m.GetAllowMultipleOccurrences())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *AttributeMappingParameterSchema) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowMultipleOccurrences sets the allowMultipleOccurrences property value. The given parameter can be provided multiple times (for example, multiple input strings in the Concatenate(string,string,...) function).
func (m *AttributeMappingParameterSchema) SetAllowMultipleOccurrences(value *bool)() {
    if m != nil {
        m.allowMultipleOccurrences = value
    }
}
// SetName sets the name property value. Parameter name.
func (m *AttributeMappingParameterSchema) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetRequired sets the required property value. true if the parameter is required; otherwise false.
func (m *AttributeMappingParameterSchema) SetRequired(value *bool)() {
    if m != nil {
        m.required = value
    }
}
// SetType sets the type property value. The type property
func (m *AttributeMappingParameterSchema) SetType(value *AttributeType)() {
    if m != nil {
        m.type_escaped = value
    }
}
