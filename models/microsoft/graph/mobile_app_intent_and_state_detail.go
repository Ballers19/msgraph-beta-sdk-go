package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MobileAppIntentAndStateDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // MobieApp identifier.
    applicationId *string;
    // The admin provided or imported title of the app.
    displayName *string;
    // Human readable version of the application
    displayVersion *string;
    // The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
    installState *ResultantAppState;
    // Mobile App Intent. Possible values are: available, notAvailable, requiredInstall, requiredUninstall, requiredAndAvailableInstall, availableInstallWithoutEnrollment, exclude.
    mobileAppIntent *MobileAppIntent;
    // The supported platforms for the app.
    supportedDeviceTypes []MobileAppSupportedDeviceType;
}
// Instantiates a new mobileAppIntentAndStateDetail and sets the default values.
func NewMobileAppIntentAndStateDetail()(*MobileAppIntentAndStateDetail) {
    m := &MobileAppIntentAndStateDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppIntentAndStateDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the applicationId property value. MobieApp identifier.
func (m *MobileAppIntentAndStateDetail) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
// Gets the displayName property value. The admin provided or imported title of the app.
func (m *MobileAppIntentAndStateDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the displayVersion property value. Human readable version of the application
func (m *MobileAppIntentAndStateDetail) GetDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayVersion
    }
}
// Gets the installState property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
func (m *MobileAppIntentAndStateDetail) GetInstallState()(*ResultantAppState) {
    if m == nil {
        return nil
    } else {
        return m.installState
    }
}
// Gets the mobileAppIntent property value. Mobile App Intent. Possible values are: available, notAvailable, requiredInstall, requiredUninstall, requiredAndAvailableInstall, availableInstallWithoutEnrollment, exclude.
func (m *MobileAppIntentAndStateDetail) GetMobileAppIntent()(*MobileAppIntent) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppIntent
    }
}
// Gets the supportedDeviceTypes property value. The supported platforms for the app.
func (m *MobileAppIntentAndStateDetail) GetSupportedDeviceTypes()([]MobileAppSupportedDeviceType) {
    if m == nil {
        return nil
    } else {
        return m.supportedDeviceTypes
    }
}
// The deserialization information for the current model
func (m *MobileAppIntentAndStateDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationId(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["displayVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayVersion(val)
        return nil
    }
    res["installState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppState)
        if err != nil {
            return err
        }
        cast := val.(ResultantAppState)
        m.SetInstallState(&cast)
        return nil
    }
    res["mobileAppIntent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppIntent)
        if err != nil {
            return err
        }
        cast := val.(MobileAppIntent)
        m.SetMobileAppIntent(&cast)
        return nil
    }
    res["supportedDeviceTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppSupportedDeviceType() })
        if err != nil {
            return err
        }
        res := make([]MobileAppSupportedDeviceType, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobileAppSupportedDeviceType))
        }
        m.SetSupportedDeviceTypes(res)
        return nil
    }
    return res
}
func (m *MobileAppIntentAndStateDetail) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MobileAppIntentAndStateDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayVersion", m.GetDisplayVersion())
        if err != nil {
            return err
        }
    }
    if m.GetInstallState() != nil {
        cast := m.GetInstallState().String()
        err := writer.WriteStringValue("installState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppIntent() != nil {
        cast := m.GetMobileAppIntent().String()
        err := writer.WriteStringValue("mobileAppIntent", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSupportedDeviceTypes()))
        for i, v := range m.GetSupportedDeviceTypes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("supportedDeviceTypes", cast)
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
func (m *MobileAppIntentAndStateDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the applicationId property value. MobieApp identifier.
// Parameters:
//  - value : Value to set for the applicationId property.
func (m *MobileAppIntentAndStateDetail) SetApplicationId(value *string)() {
    m.applicationId = value
}
// Sets the displayName property value. The admin provided or imported title of the app.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *MobileAppIntentAndStateDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the displayVersion property value. Human readable version of the application
// Parameters:
//  - value : Value to set for the displayVersion property.
func (m *MobileAppIntentAndStateDetail) SetDisplayVersion(value *string)() {
    m.displayVersion = value
}
// Sets the installState property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
// Parameters:
//  - value : Value to set for the installState property.
func (m *MobileAppIntentAndStateDetail) SetInstallState(value *ResultantAppState)() {
    m.installState = value
}
// Sets the mobileAppIntent property value. Mobile App Intent. Possible values are: available, notAvailable, requiredInstall, requiredUninstall, requiredAndAvailableInstall, availableInstallWithoutEnrollment, exclude.
// Parameters:
//  - value : Value to set for the mobileAppIntent property.
func (m *MobileAppIntentAndStateDetail) SetMobileAppIntent(value *MobileAppIntent)() {
    m.mobileAppIntent = value
}
// Sets the supportedDeviceTypes property value. The supported platforms for the app.
// Parameters:
//  - value : Value to set for the supportedDeviceTypes property.
func (m *MobileAppIntentAndStateDetail) SetSupportedDeviceTypes(value []MobileAppSupportedDeviceType)() {
    m.supportedDeviceTypes = value
}
