package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingApplicabilityable 
type DeviceManagementConfigurationSettingApplicabilityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDeviceMode()(*DeviceManagementConfigurationDeviceMode)
    GetPlatform()(*DeviceManagementConfigurationPlatforms)
    GetTechnologies()(*DeviceManagementConfigurationTechnologies)
    GetType()(*string)
    SetDescription(value *string)()
    SetDeviceMode(value *DeviceManagementConfigurationDeviceMode)()
    SetPlatform(value *DeviceManagementConfigurationPlatforms)()
    SetTechnologies(value *DeviceManagementConfigurationTechnologies)()
    SetType(value *string)()
}
