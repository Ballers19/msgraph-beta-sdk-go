package directoryroles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i28acc41a9ddaa5012f518de9ebb17e81fc5d01a5bb443c4043b0cabb5ae0b5ac "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/count"
    i686e52110d509abde81e2ec02a4b1cacb26be255a28efb139b5fe07005585dcb "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/getuserownedobjects"
    i6c75901f09b564ec245b7b3a868272e273083db1b9879e0b05d7329117fa6e34 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/getbyids"
    i89f3d18e33abe74e12cff2eb112be1d34453a000675f0e7670b07807419b57ff "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/delta"
    ifd6f8f935853943530de1cc46edc30d54c90acd20d9483c7525660fa0c70d33f "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/validateproperties"
)

// DirectoryRolesRequestBuilder provides operations to manage the collection of directoryRole entities.
type DirectoryRolesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryRolesRequestBuilderGetQueryParameters list the directory roles that are activated in the tenant. This operation only returns roles that have been activated. A role becomes activated when an admin activates the role using the Activate directoryRole API. Not all built-in roles are initially activated.  When assigning a role using the Azure portal, the role activation step is implicitly done on the admin's behalf. To get the full list of roles that are available in Azure AD, use List directoryRoleTemplates.
type DirectoryRolesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
}
// DirectoryRolesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRolesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRolesRequestBuilderGetQueryParameters
}
// DirectoryRolesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRolesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDirectoryRolesRequestBuilderInternal instantiates a new DirectoryRolesRequestBuilder and sets the default values.
func NewDirectoryRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    m := &DirectoryRolesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles{?%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRolesRequestBuilder instantiates a new DirectoryRolesRequestBuilder and sets the default values.
func NewDirectoryRolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRolesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *DirectoryRolesRequestBuilder) Count()(*i28acc41a9ddaa5012f518de9ebb17e81fc5d01a5bb443c4043b0cabb5ae0b5ac.CountRequestBuilder) {
    return i28acc41a9ddaa5012f518de9ebb17e81fc5d01a5bb443c4043b0cabb5ae0b5ac.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation list the directory roles that are activated in the tenant. This operation only returns roles that have been activated. A role becomes activated when an admin activates the role using the Activate directoryRole API. Not all built-in roles are initially activated.  When assigning a role using the Azure portal, the role activation step is implicitly done on the admin's behalf. To get the full list of roles that are available in Azure AD, use List directoryRoleTemplates.
func (m *DirectoryRolesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration list the directory roles that are activated in the tenant. This operation only returns roles that have been activated. A role becomes activated when an admin activates the role using the Activate directoryRole API. Not all built-in roles are initially activated.  When assigning a role using the Azure portal, the role activation step is implicitly done on the admin's behalf. To get the full list of roles that are available in Azure AD, use List directoryRoleTemplates.
func (m *DirectoryRolesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DirectoryRolesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformation activate a directory role. To read a directory role or update its members, it must first be activated in the tenant. Only the Company Administrators  and the implicit Users directory roles are activated by default. To access and assign members to another directory role, you must first activate it with its corresponding directory role template (directoryRoleTemplate).
func (m *DirectoryRolesRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration activate a directory role. To read a directory role or update its members, it must first be activated in the tenant. Only the Company Administrators  and the implicit Users directory roles are activated by default. To access and assign members to another directory role, you must first activate it with its corresponding directory role template (directoryRoleTemplate).
func (m *DirectoryRolesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable, requestConfiguration *DirectoryRolesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *DirectoryRolesRequestBuilder) Delta()(*i89f3d18e33abe74e12cff2eb112be1d34453a000675f0e7670b07807419b57ff.DeltaRequestBuilder) {
    return i89f3d18e33abe74e12cff2eb112be1d34453a000675f0e7670b07807419b57ff.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get list the directory roles that are activated in the tenant. This operation only returns roles that have been activated. A role becomes activated when an admin activates the role using the Activate directoryRole API. Not all built-in roles are initially activated.  When assigning a role using the Azure portal, the role activation step is implicitly done on the admin's behalf. To get the full list of roles that are available in Azure AD, use List directoryRoleTemplates.
func (m *DirectoryRolesRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetByIds the getByIds property
func (m *DirectoryRolesRequestBuilder) GetByIds()(*i6c75901f09b564ec245b7b3a868272e273083db1b9879e0b05d7329117fa6e34.GetByIdsRequestBuilder) {
    return i6c75901f09b564ec245b7b3a868272e273083db1b9879e0b05d7329117fa6e34.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects the getUserOwnedObjects property
func (m *DirectoryRolesRequestBuilder) GetUserOwnedObjects()(*i686e52110d509abde81e2ec02a4b1cacb26be255a28efb139b5fe07005585dcb.GetUserOwnedObjectsRequestBuilder) {
    return i686e52110d509abde81e2ec02a4b1cacb26be255a28efb139b5fe07005585dcb.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler list the directory roles that are activated in the tenant. This operation only returns roles that have been activated. A role becomes activated when an admin activates the role using the Activate directoryRole API. Not all built-in roles are initially activated.  When assigning a role using the Azure portal, the role activation step is implicitly done on the admin's behalf. To get the full list of roles that are available in Azure AD, use List directoryRoleTemplates.
func (m *DirectoryRolesRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DirectoryRolesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryRoleCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleCollectionResponseable), nil
}
// Post activate a directory role. To read a directory role or update its members, it must first be activated in the tenant. Only the Company Administrators  and the implicit Users directory roles are activated by default. To access and assign members to another directory role, you must first activate it with its corresponding directory role template (directoryRoleTemplate).
func (m *DirectoryRolesRequestBuilder) Post(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler activate a directory role. To read a directory role or update its members, it must first be activated in the tenant. Only the Company Administrators  and the implicit Users directory roles are activated by default. To access and assign members to another directory role, you must first activate it with its corresponding directory role template (directoryRoleTemplate).
func (m *DirectoryRolesRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable, requestConfiguration *DirectoryRolesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryRoleFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable), nil
}
// ValidateProperties the validateProperties property
func (m *DirectoryRolesRequestBuilder) ValidateProperties()(*ifd6f8f935853943530de1cc46edc30d54c90acd20d9483c7525660fa0c70d33f.ValidatePropertiesRequestBuilder) {
    return ifd6f8f935853943530de1cc46edc30d54c90acd20d9483c7525660fa0c70d33f.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
