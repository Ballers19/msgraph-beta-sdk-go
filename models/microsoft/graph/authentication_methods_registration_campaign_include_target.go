package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AuthenticationMethodsRegistrationCampaignIncludeTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The object identifier of an Azure AD user or group.
    id *string;
    // The authentication method that the user is prompted to register. The value must be microsoftAuthenticator.
    targetedAuthenticationMethod *string;
    // The type of the authentication method target. Possible values are: user, group, unknownFutureValue.
    targetType *AuthenticationMethodTargetType;
}
// Instantiates a new authenticationMethodsRegistrationCampaignIncludeTarget and sets the default values.
func NewAuthenticationMethodsRegistrationCampaignIncludeTarget()(*AuthenticationMethodsRegistrationCampaignIncludeTarget) {
    m := &AuthenticationMethodsRegistrationCampaignIncludeTarget{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the id property value. The object identifier of an Azure AD user or group.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the targetedAuthenticationMethod property value. The authentication method that the user is prompted to register. The value must be microsoftAuthenticator.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetTargetedAuthenticationMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetedAuthenticationMethod
    }
}
// Gets the targetType property value. The type of the authentication method target. Possible values are: user, group, unknownFutureValue.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetTargetType()(*AuthenticationMethodTargetType) {
    if m == nil {
        return nil
    } else {
        return m.targetType
    }
}
// The deserialization information for the current model
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["targetedAuthenticationMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetedAuthenticationMethod(val)
        return nil
    }
    res["targetType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodTargetType)
        if err != nil {
            return err
        }
        cast := val.(AuthenticationMethodTargetType)
        m.SetTargetType(&cast)
        return nil
    }
    return res
}
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetedAuthenticationMethod", m.GetTargetedAuthenticationMethod())
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := m.GetTargetType().String()
        err := writer.WriteStringValue("targetType", &cast)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the id property value. The object identifier of an Azure AD user or group.
// Parameters:
//  - value : Value to set for the id property.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetId(value *string)() {
    m.id = value
}
// Sets the targetedAuthenticationMethod property value. The authentication method that the user is prompted to register. The value must be microsoftAuthenticator.
// Parameters:
//  - value : Value to set for the targetedAuthenticationMethod property.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetTargetedAuthenticationMethod(value *string)() {
    m.targetedAuthenticationMethod = value
}
// Sets the targetType property value. The type of the authentication method target. Possible values are: user, group, unknownFutureValue.
// Parameters:
//  - value : Value to set for the targetType property.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetTargetType(value *AuthenticationMethodTargetType)() {
    m.targetType = value
}
