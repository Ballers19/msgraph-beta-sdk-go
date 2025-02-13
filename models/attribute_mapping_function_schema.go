package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeMappingFunctionSchema 
type AttributeMappingFunctionSchema struct {
    Entity
    // Collection of function parameters.
    parameters []AttributeMappingParameterSchemaable
}
// NewAttributeMappingFunctionSchema instantiates a new attributeMappingFunctionSchema and sets the default values.
func NewAttributeMappingFunctionSchema()(*AttributeMappingFunctionSchema) {
    m := &AttributeMappingFunctionSchema{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAttributeMappingFunctionSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeMappingFunctionSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeMappingFunctionSchema(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMappingFunctionSchema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttributeMappingParameterSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeMappingParameterSchemaable, len(val))
            for i, v := range val {
                res[i] = v.(AttributeMappingParameterSchemaable)
            }
            m.SetParameters(res)
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. Collection of function parameters.
func (m *AttributeMappingFunctionSchema) GetParameters()([]AttributeMappingParameterSchemaable) {
    if m == nil {
        return nil
    } else {
        return m.parameters
    }
}
// Serialize serializes information the current object
func (m *AttributeMappingFunctionSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetParameters sets the parameters property value. Collection of function parameters.
func (m *AttributeMappingFunctionSchema) SetParameters(value []AttributeMappingParameterSchemaable)() {
    if m != nil {
        m.parameters = value
    }
}
