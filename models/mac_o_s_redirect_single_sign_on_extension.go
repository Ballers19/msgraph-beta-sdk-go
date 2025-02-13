package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSRedirectSingleSignOnExtension 
type MacOSRedirectSingleSignOnExtension struct {
    MacOSSingleSignOnExtension
    // Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
    configurations []KeyTypedValuePairable
    // Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
    extensionIdentifier *string
    // Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
    teamIdentifier *string
    // One or more URL prefixes of identity providers on whose behalf the app extension performs single sign-on. URLs must begin with http:// or https://. All URL prefixes must be unique for all profiles.
    urlPrefixes []string
}
// NewMacOSRedirectSingleSignOnExtension instantiates a new MacOSRedirectSingleSignOnExtension and sets the default values.
func NewMacOSRedirectSingleSignOnExtension()(*MacOSRedirectSingleSignOnExtension) {
    m := &MacOSRedirectSingleSignOnExtension{
        MacOSSingleSignOnExtension: *NewMacOSSingleSignOnExtension(),
    }
    return m
}
// CreateMacOSRedirectSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSRedirectSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSRedirectSingleSignOnExtension(), nil
}
// GetConfigurations gets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *MacOSRedirectSingleSignOnExtension) GetConfigurations()([]KeyTypedValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.configurations
    }
}
// GetExtensionIdentifier gets the extensionIdentifier property value. Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
func (m *MacOSRedirectSingleSignOnExtension) GetExtensionIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionIdentifier
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSRedirectSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MacOSSingleSignOnExtension.GetFieldDeserializers()
    res["configurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyTypedValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyTypedValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyTypedValuePairable)
            }
            m.SetConfigurations(res)
        }
        return nil
    }
    res["extensionIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionIdentifier(val)
        }
        return nil
    }
    res["teamIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamIdentifier(val)
        }
        return nil
    }
    res["urlPrefixes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUrlPrefixes(res)
        }
        return nil
    }
    return res
}
// GetTeamIdentifier gets the teamIdentifier property value. Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
func (m *MacOSRedirectSingleSignOnExtension) GetTeamIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamIdentifier
    }
}
// GetUrlPrefixes gets the urlPrefixes property value. One or more URL prefixes of identity providers on whose behalf the app extension performs single sign-on. URLs must begin with http:// or https://. All URL prefixes must be unique for all profiles.
func (m *MacOSRedirectSingleSignOnExtension) GetUrlPrefixes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.urlPrefixes
    }
}
// Serialize serializes information the current object
func (m *MacOSRedirectSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MacOSSingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurations()))
        for i, v := range m.GetConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("configurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("extensionIdentifier", m.GetExtensionIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamIdentifier", m.GetTeamIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetUrlPrefixes() != nil {
        err = writer.WriteCollectionOfStringValues("urlPrefixes", m.GetUrlPrefixes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurations sets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *MacOSRedirectSingleSignOnExtension) SetConfigurations(value []KeyTypedValuePairable)() {
    if m != nil {
        m.configurations = value
    }
}
// SetExtensionIdentifier sets the extensionIdentifier property value. Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
func (m *MacOSRedirectSingleSignOnExtension) SetExtensionIdentifier(value *string)() {
    if m != nil {
        m.extensionIdentifier = value
    }
}
// SetTeamIdentifier sets the teamIdentifier property value. Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
func (m *MacOSRedirectSingleSignOnExtension) SetTeamIdentifier(value *string)() {
    if m != nil {
        m.teamIdentifier = value
    }
}
// SetUrlPrefixes sets the urlPrefixes property value. One or more URL prefixes of identity providers on whose behalf the app extension performs single sign-on. URLs must begin with http:// or https://. All URL prefixes must be unique for all profiles.
func (m *MacOSRedirectSingleSignOnExtension) SetUrlPrefixes(value []string)() {
    if m != nil {
        m.urlPrefixes = value
    }
}
