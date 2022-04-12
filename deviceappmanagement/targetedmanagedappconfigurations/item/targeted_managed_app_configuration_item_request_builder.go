package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i111455d8fd8d7f0ddf817ca0fcc2bc0624e3819c5bdee6ec2194f7ae9db42ddf "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/targetapps"
    i4633bb1fbea3b9ff62e44e3637dd657b2fff6cbd90287bd10a013fed2d40860f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments"
    i7d0f522bac2e00da7e218c106fc841a9805ca78f7c4bdd7512999c9cfc802d42 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps"
    ida25c58687771cbc2e8cd5750db08465dd94e755be0fd37d473f36b0a0502e49 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/deploymentsummary"
    iefd9a3454d070af5a6b3c6f36a2215b5d86ef0d740bac940405f70157b57b8a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assign"
    i6a789da51c915c8f551873e29bac081934811f780b24231053b29a8984f30fb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps/item"
    ie924ef4384ce239f82e862609ff7154fd7dc52ac669aa7e01fb54a9dbd903a66 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments/item"
)

// TargetedManagedAppConfigurationItemRequestBuilder provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
type TargetedManagedAppConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions options for Delete
type TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// TargetedManagedAppConfigurationItemRequestBuilderGetOptions options for Get
type TargetedManagedAppConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// TargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters targeted managed app configurations.
type TargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TargetedManagedAppConfigurationItemRequestBuilderPatchOptions options for Patch
type TargetedManagedAppConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppConfigurationable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Apps the apps property
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Apps()(*i7d0f522bac2e00da7e218c106fc841a9805ca78f7c4bdd7512999c9cfc802d42.AppsRequestBuilder) {
    return i7d0f522bac2e00da7e218c106fc841a9805ca78f7c4bdd7512999c9cfc802d42.NewAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item.apps.item collection
func (m *TargetedManagedAppConfigurationItemRequestBuilder) AppsById(id string)(*i6a789da51c915c8f551873e29bac081934811f780b24231053b29a8984f30fb5.ManagedMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp%2Did"] = id
    }
    return i6a789da51c915c8f551873e29bac081934811f780b24231053b29a8984f30fb5.NewManagedMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Assign the assign property
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Assign()(*iefd9a3454d070af5a6b3c6f36a2215b5d86ef0d740bac940405f70157b57b8a5.AssignRequestBuilder) {
    return iefd9a3454d070af5a6b3c6f36a2215b5d86ef0d740bac940405f70157b57b8a5.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Assignments()(*i4633bb1fbea3b9ff62e44e3637dd657b2fff6cbd90287bd10a013fed2d40860f.AssignmentsRequestBuilder) {
    return i4633bb1fbea3b9ff62e44e3637dd657b2fff6cbd90287bd10a013fed2d40860f.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item.assignments.item collection
func (m *TargetedManagedAppConfigurationItemRequestBuilder) AssignmentsById(id string)(*ie924ef4384ce239f82e862609ff7154fd7dc52ac669aa7e01fb54a9dbd903a66.TargetedManagedAppPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppPolicyAssignment%2Did"] = id
    }
    return ie924ef4384ce239f82e862609ff7154fd7dc52ac669aa7e01fb54a9dbd903a66.NewTargetedManagedAppPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTargetedManagedAppConfigurationItemRequestBuilderInternal instantiates a new TargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppConfigurationItemRequestBuilder) {
    m := &TargetedManagedAppConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTargetedManagedAppConfigurationItemRequestBuilder instantiates a new TargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property targetedManagedAppConfigurations for deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation targeted managed app configurations.
func (m *TargetedManagedAppConfigurationItemRequestBuilder) CreateGetRequestInformation(options *TargetedManagedAppConfigurationItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property targetedManagedAppConfigurations in deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *TargetedManagedAppConfigurationItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property targetedManagedAppConfigurations for deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Delete(options *TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeploymentSummary the deploymentSummary property
func (m *TargetedManagedAppConfigurationItemRequestBuilder) DeploymentSummary()(*ida25c58687771cbc2e8cd5750db08465dd94e755be0fd37d473f36b0a0502e49.DeploymentSummaryRequestBuilder) {
    return ida25c58687771cbc2e8cd5750db08465dd94e755be0fd37d473f36b0a0502e49.NewDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get targeted managed app configurations.
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Get(options *TargetedManagedAppConfigurationItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTargetedManagedAppConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppConfigurationable), nil
}
// Patch update the navigation property targetedManagedAppConfigurations in deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Patch(options *TargetedManagedAppConfigurationItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// TargetApps the targetApps property
func (m *TargetedManagedAppConfigurationItemRequestBuilder) TargetApps()(*i111455d8fd8d7f0ddf817ca0fcc2bc0624e3819c5bdee6ec2194f7ae9db42ddf.TargetAppsRequestBuilder) {
    return i111455d8fd8d7f0ddf817ca0fcc2bc0624e3819c5bdee6ec2194f7ae9db42ddf.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
