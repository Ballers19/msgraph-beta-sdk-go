package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationConnectionSettings 
type EducationSynchronizationConnectionSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Client ID used to connect to the provider.
    clientId *string
    // Client secret to authenticate the connection to the provider.
    clientSecret *string
    // The type property
    type_escaped *string
}
// NewEducationSynchronizationConnectionSettings instantiates a new educationSynchronizationConnectionSettings and sets the default values.
func NewEducationSynchronizationConnectionSettings()(*EducationSynchronizationConnectionSettings) {
    m := &EducationSynchronizationConnectionSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odatatypeValue := "#microsoft.graph.educationSynchronizationConnectionSettings";
    m.SetType(&odatatypeValue);
    return m
}
// CreateEducationSynchronizationConnectionSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationConnectionSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.educationSynchronizationOAuth1ConnectionSettings":
                        return NewEducationSynchronizationOAuth1ConnectionSettings(), nil
                    case "#microsoft.graph.educationSynchronizationOAuth2ClientCredentialsConnectionSettings":
                        return NewEducationSynchronizationOAuth2ClientCredentialsConnectionSettings(), nil
                }
            }
        }
    }
    return NewEducationSynchronizationConnectionSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationSynchronizationConnectionSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClientId gets the clientId property value. Client ID used to connect to the provider.
func (m *EducationSynchronizationConnectionSettings) GetClientId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientId
    }
}
// GetClientSecret gets the clientSecret property value. Client secret to authenticate the connection to the provider.
func (m *EducationSynchronizationConnectionSettings) GetClientSecret()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientSecret
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationConnectionSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["clientSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetType gets the @odata.type property value. The type property
func (m *EducationSynchronizationConnectionSettings) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *EducationSynchronizationConnectionSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientSecret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetType())
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
func (m *EducationSynchronizationConnectionSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClientId sets the clientId property value. Client ID used to connect to the provider.
func (m *EducationSynchronizationConnectionSettings) SetClientId(value *string)() {
    if m != nil {
        m.clientId = value
    }
}
// SetClientSecret sets the clientSecret property value. Client secret to authenticate the connection to the provider.
func (m *EducationSynchronizationConnectionSettings) SetClientSecret(value *string)() {
    if m != nil {
        m.clientSecret = value
    }
}
// SetType sets the @odata.type property value. The type property
func (m *EducationSynchronizationConnectionSettings) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
