package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PhoneAuthenticationMethod struct {
    AuthenticationMethod
    // The phone number to text or call for authentication. Phone numbers use the format '+<country code> <number>x<extension>', with extension optional. For example, +1 5555551234 or +1 5555551234x123 are valid. Numbers are rejected when creating/updating if they do not match the required format.
    phoneNumber *string;
    // The type of this phone. Possible values are: mobile, alternateMobile, or office.
    phoneType *AuthenticationPhoneType;
    // Whether a phone is ready to be used for SMS sign-in or not. Possible values are: notSupported, notAllowedByPolicy, notEnabled, phoneNumberNotUnique, ready, or notConfigured, unknownFutureValue.
    smsSignInState *AuthenticationMethodSignInState;
}
// Instantiates a new phoneAuthenticationMethod and sets the default values.
func NewPhoneAuthenticationMethod()(*PhoneAuthenticationMethod) {
    m := &PhoneAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// Gets the phoneNumber property value. The phone number to text or call for authentication. Phone numbers use the format '+<country code> <number>x<extension>', with extension optional. For example, +1 5555551234 or +1 5555551234x123 are valid. Numbers are rejected when creating/updating if they do not match the required format.
func (m *PhoneAuthenticationMethod) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// Gets the phoneType property value. The type of this phone. Possible values are: mobile, alternateMobile, or office.
func (m *PhoneAuthenticationMethod) GetPhoneType()(*AuthenticationPhoneType) {
    if m == nil {
        return nil
    } else {
        return m.phoneType
    }
}
// Gets the smsSignInState property value. Whether a phone is ready to be used for SMS sign-in or not. Possible values are: notSupported, notAllowedByPolicy, notEnabled, phoneNumberNotUnique, ready, or notConfigured, unknownFutureValue.
func (m *PhoneAuthenticationMethod) GetSmsSignInState()(*AuthenticationMethodSignInState) {
    if m == nil {
        return nil
    } else {
        return m.smsSignInState
    }
}
// The deserialization information for the current model
func (m *PhoneAuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPhoneNumber(val)
        return nil
    }
    res["phoneType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationPhoneType)
        if err != nil {
            return err
        }
        cast := val.(AuthenticationPhoneType)
        m.SetPhoneType(&cast)
        return nil
    }
    res["smsSignInState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodSignInState)
        if err != nil {
            return err
        }
        cast := val.(AuthenticationMethodSignInState)
        m.SetSmsSignInState(&cast)
        return nil
    }
    return res
}
func (m *PhoneAuthenticationMethod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PhoneAuthenticationMethod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    if m.GetPhoneType() != nil {
        cast := m.GetPhoneType().String()
        err = writer.WriteStringValue("phoneType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSmsSignInState() != nil {
        cast := m.GetSmsSignInState().String()
        err = writer.WriteStringValue("smsSignInState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the phoneNumber property value. The phone number to text or call for authentication. Phone numbers use the format '+<country code> <number>x<extension>', with extension optional. For example, +1 5555551234 or +1 5555551234x123 are valid. Numbers are rejected when creating/updating if they do not match the required format.
// Parameters:
//  - value : Value to set for the phoneNumber property.
func (m *PhoneAuthenticationMethod) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// Sets the phoneType property value. The type of this phone. Possible values are: mobile, alternateMobile, or office.
// Parameters:
//  - value : Value to set for the phoneType property.
func (m *PhoneAuthenticationMethod) SetPhoneType(value *AuthenticationPhoneType)() {
    m.phoneType = value
}
// Sets the smsSignInState property value. Whether a phone is ready to be used for SMS sign-in or not. Possible values are: notSupported, notAllowedByPolicy, notEnabled, phoneNumberNotUnique, ready, or notConfigured, unknownFutureValue.
// Parameters:
//  - value : Value to set for the smsSignInState property.
func (m *PhoneAuthenticationMethod) SetSmsSignInState(value *AuthenticationMethodSignInState)() {
    m.smsSignInState = value
}
