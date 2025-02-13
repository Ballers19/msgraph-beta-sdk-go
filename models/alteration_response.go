package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlterationResponse 
type AlterationResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Defines the original user query string.
    originalQueryString *string
    // Defines the details of alteration information for the spelling correction.
    queryAlteration SearchAlterationable
    // Defines the type of the spelling correction. Possible values are suggestion, modification.
    queryAlterationType *SearchAlterationType
}
// NewAlterationResponse instantiates a new alterationResponse and sets the default values.
func NewAlterationResponse()(*AlterationResponse) {
    m := &AlterationResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAlterationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlterationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlterationResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlterationResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlterationResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["originalQueryString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalQueryString(val)
        }
        return nil
    }
    res["queryAlteration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchAlterationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryAlteration(val.(SearchAlterationable))
        }
        return nil
    }
    res["queryAlterationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSearchAlterationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryAlterationType(val.(*SearchAlterationType))
        }
        return nil
    }
    return res
}
// GetOriginalQueryString gets the originalQueryString property value. Defines the original user query string.
func (m *AlterationResponse) GetOriginalQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originalQueryString
    }
}
// GetQueryAlteration gets the queryAlteration property value. Defines the details of alteration information for the spelling correction.
func (m *AlterationResponse) GetQueryAlteration()(SearchAlterationable) {
    if m == nil {
        return nil
    } else {
        return m.queryAlteration
    }
}
// GetQueryAlterationType gets the queryAlterationType property value. Defines the type of the spelling correction. Possible values are suggestion, modification.
func (m *AlterationResponse) GetQueryAlterationType()(*SearchAlterationType) {
    if m == nil {
        return nil
    } else {
        return m.queryAlterationType
    }
}
// Serialize serializes information the current object
func (m *AlterationResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("originalQueryString", m.GetOriginalQueryString())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("queryAlteration", m.GetQueryAlteration())
        if err != nil {
            return err
        }
    }
    if m.GetQueryAlterationType() != nil {
        cast := (*m.GetQueryAlterationType()).String()
        err := writer.WriteStringValue("queryAlterationType", &cast)
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
func (m *AlterationResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOriginalQueryString sets the originalQueryString property value. Defines the original user query string.
func (m *AlterationResponse) SetOriginalQueryString(value *string)() {
    if m != nil {
        m.originalQueryString = value
    }
}
// SetQueryAlteration sets the queryAlteration property value. Defines the details of alteration information for the spelling correction.
func (m *AlterationResponse) SetQueryAlteration(value SearchAlterationable)() {
    if m != nil {
        m.queryAlteration = value
    }
}
// SetQueryAlterationType sets the queryAlterationType property value. Defines the type of the spelling correction. Possible values are suggestion, modification.
func (m *AlterationResponse) SetQueryAlterationType(value *SearchAlterationType)() {
    if m != nil {
        m.queryAlterationType = value
    }
}
