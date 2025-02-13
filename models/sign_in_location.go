package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SignInLocation 
type SignInLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Provides the city where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
    city *string
    // Provides the country code info (2 letter code) where the sign-in originated.  This is calculated using latitude/longitude information from the sign-in activity.
    countryOrRegion *string
    // Provides the latitude, longitude and altitude where the sign-in originated.
    geoCoordinates GeoCoordinatesable
    // Provides the State where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
    state *string
}
// NewSignInLocation instantiates a new signInLocation and sets the default values.
func NewSignInLocation()(*SignInLocation) {
    m := &SignInLocation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSignInLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSignInLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSignInLocation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SignInLocation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCity gets the city property value. Provides the city where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
func (m *SignInLocation) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// GetCountryOrRegion gets the countryOrRegion property value. Provides the country code info (2 letter code) where the sign-in originated.  This is calculated using latitude/longitude information from the sign-in activity.
func (m *SignInLocation) GetCountryOrRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SignInLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["countryOrRegion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryOrRegion(val)
        }
        return nil
    }
    res["geoCoordinates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGeoCoordinatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeoCoordinates(val.(GeoCoordinatesable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
// GetGeoCoordinates gets the geoCoordinates property value. Provides the latitude, longitude and altitude where the sign-in originated.
func (m *SignInLocation) GetGeoCoordinates()(GeoCoordinatesable) {
    if m == nil {
        return nil
    } else {
        return m.geoCoordinates
    }
}
// GetState gets the state property value. Provides the State where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
func (m *SignInLocation) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Serialize serializes information the current object
func (m *SignInLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("countryOrRegion", m.GetCountryOrRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("geoCoordinates", m.GetGeoCoordinates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *SignInLocation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCity sets the city property value. Provides the city where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
func (m *SignInLocation) SetCity(value *string)() {
    if m != nil {
        m.city = value
    }
}
// SetCountryOrRegion sets the countryOrRegion property value. Provides the country code info (2 letter code) where the sign-in originated.  This is calculated using latitude/longitude information from the sign-in activity.
func (m *SignInLocation) SetCountryOrRegion(value *string)() {
    if m != nil {
        m.countryOrRegion = value
    }
}
// SetGeoCoordinates sets the geoCoordinates property value. Provides the latitude, longitude and altitude where the sign-in originated.
func (m *SignInLocation) SetGeoCoordinates(value GeoCoordinatesable)() {
    if m != nil {
        m.geoCoordinates = value
    }
}
// SetState sets the state property value. Provides the State where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity.
func (m *SignInLocation) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
