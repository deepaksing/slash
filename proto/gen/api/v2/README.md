# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v2/common.proto](#api_v2_common-proto)
    - [RowStatus](#slash-api-v2-RowStatus)
  
- [api/v2/shortcut_service.proto](#api_v2_shortcut_service-proto)
    - [CreateShortcutRequest](#slash-api-v2-CreateShortcutRequest)
    - [CreateShortcutResponse](#slash-api-v2-CreateShortcutResponse)
    - [GetShortcutRequest](#slash-api-v2-GetShortcutRequest)
    - [GetShortcutResponse](#slash-api-v2-GetShortcutResponse)
    - [ListShortcutsRequest](#slash-api-v2-ListShortcutsRequest)
    - [ListShortcutsResponse](#slash-api-v2-ListShortcutsResponse)
    - [OpenGraphMetadata](#slash-api-v2-OpenGraphMetadata)
    - [Shortcut](#slash-api-v2-Shortcut)
  
    - [Visibility](#slash-api-v2-Visibility)
  
    - [ShortcutService](#slash-api-v2-ShortcutService)
  
- [api/v2/user_service.proto](#api_v2_user_service-proto)
    - [CreateUserAccessTokenRequest](#slash-api-v2-CreateUserAccessTokenRequest)
    - [CreateUserAccessTokenResponse](#slash-api-v2-CreateUserAccessTokenResponse)
    - [CreateUserRequest](#slash-api-v2-CreateUserRequest)
    - [CreateUserResponse](#slash-api-v2-CreateUserResponse)
    - [DeleteUserAccessTokenRequest](#slash-api-v2-DeleteUserAccessTokenRequest)
    - [DeleteUserAccessTokenResponse](#slash-api-v2-DeleteUserAccessTokenResponse)
    - [DeleteUserRequest](#slash-api-v2-DeleteUserRequest)
    - [DeleteUserResponse](#slash-api-v2-DeleteUserResponse)
    - [GetUserRequest](#slash-api-v2-GetUserRequest)
    - [GetUserResponse](#slash-api-v2-GetUserResponse)
    - [ListUserAccessTokensRequest](#slash-api-v2-ListUserAccessTokensRequest)
    - [ListUserAccessTokensResponse](#slash-api-v2-ListUserAccessTokensResponse)
    - [ListUsersRequest](#slash-api-v2-ListUsersRequest)
    - [ListUsersResponse](#slash-api-v2-ListUsersResponse)
    - [User](#slash-api-v2-User)
    - [UserAccessToken](#slash-api-v2-UserAccessToken)
  
    - [Role](#slash-api-v2-Role)
  
    - [UserService](#slash-api-v2-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v2_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/common.proto


 


<a name="slash-api-v2-RowStatus"></a>

### RowStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| ROW_STATUS_UNSPECIFIED | 0 |  |
| NORMAL | 1 |  |
| ARCHIVED | 2 |  |


 

 

 



<a name="api_v2_shortcut_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/shortcut_service.proto



<a name="slash-api-v2-CreateShortcutRequest"></a>

### CreateShortcutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shortcut | [Shortcut](#slash-api-v2-Shortcut) |  |  |






<a name="slash-api-v2-CreateShortcutResponse"></a>

### CreateShortcutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shortcut | [Shortcut](#slash-api-v2-Shortcut) |  |  |






<a name="slash-api-v2-GetShortcutRequest"></a>

### GetShortcutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="slash-api-v2-GetShortcutResponse"></a>

### GetShortcutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shortcut | [Shortcut](#slash-api-v2-Shortcut) |  |  |






<a name="slash-api-v2-ListShortcutsRequest"></a>

### ListShortcutsRequest







<a name="slash-api-v2-ListShortcutsResponse"></a>

### ListShortcutsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shortcuts | [Shortcut](#slash-api-v2-Shortcut) | repeated |  |






<a name="slash-api-v2-OpenGraphMetadata"></a>

### OpenGraphMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| image | [string](#string) |  |  |






<a name="slash-api-v2-Shortcut"></a>

### Shortcut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| creator_id | [int32](#int32) |  |  |
| created_ts | [int64](#int64) |  |  |
| updated_ts | [int64](#int64) |  |  |
| row_status | [RowStatus](#slash-api-v2-RowStatus) |  |  |
| name | [string](#string) |  |  |
| link | [string](#string) |  |  |
| title | [string](#string) |  |  |
| tags | [string](#string) | repeated |  |
| description | [string](#string) |  |  |
| visibility | [Visibility](#slash-api-v2-Visibility) |  |  |
| og_metadata | [OpenGraphMetadata](#slash-api-v2-OpenGraphMetadata) |  |  |





 


<a name="slash-api-v2-Visibility"></a>

### Visibility


| Name | Number | Description |
| ---- | ------ | ----------- |
| VISIBILITY_UNSPECIFIED | 0 |  |
| PRIVATE | 1 |  |
| WORKSPACE | 2 |  |
| PUBLIC | 3 |  |


 

 


<a name="slash-api-v2-ShortcutService"></a>

### ShortcutService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListShortcuts | [ListShortcutsRequest](#slash-api-v2-ListShortcutsRequest) | [ListShortcutsResponse](#slash-api-v2-ListShortcutsResponse) | ListShortcuts returns a list of shortcuts. |
| GetShortcut | [GetShortcutRequest](#slash-api-v2-GetShortcutRequest) | [GetShortcutResponse](#slash-api-v2-GetShortcutResponse) | GetShortcut returns a shortcut by name. |
| CreateShortcut | [CreateShortcutRequest](#slash-api-v2-CreateShortcutRequest) | [CreateShortcutResponse](#slash-api-v2-CreateShortcutResponse) | CreateShortcut creates a shortcut. |

 



<a name="api_v2_user_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/user_service.proto



<a name="slash-api-v2-CreateUserAccessTokenRequest"></a>

### CreateUserAccessTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | id is the user id. |
| user_access_token | [UserAccessToken](#slash-api-v2-UserAccessToken) |  |  |






<a name="slash-api-v2-CreateUserAccessTokenResponse"></a>

### CreateUserAccessTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [UserAccessToken](#slash-api-v2-UserAccessToken) |  |  |






<a name="slash-api-v2-CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#slash-api-v2-User) |  |  |






<a name="slash-api-v2-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#slash-api-v2-User) |  |  |






<a name="slash-api-v2-DeleteUserAccessTokenRequest"></a>

### DeleteUserAccessTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | id is the user id. |
| access_token | [string](#string) |  | access_token is the access token to delete. |






<a name="slash-api-v2-DeleteUserAccessTokenResponse"></a>

### DeleteUserAccessTokenResponse







<a name="slash-api-v2-DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="slash-api-v2-DeleteUserResponse"></a>

### DeleteUserResponse







<a name="slash-api-v2-GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="slash-api-v2-GetUserResponse"></a>

### GetUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#slash-api-v2-User) |  |  |






<a name="slash-api-v2-ListUserAccessTokensRequest"></a>

### ListUserAccessTokensRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | id is the user id. |






<a name="slash-api-v2-ListUserAccessTokensResponse"></a>

### ListUserAccessTokensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_tokens | [UserAccessToken](#slash-api-v2-UserAccessToken) | repeated |  |






<a name="slash-api-v2-ListUsersRequest"></a>

### ListUsersRequest







<a name="slash-api-v2-ListUsersResponse"></a>

### ListUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#slash-api-v2-User) | repeated |  |






<a name="slash-api-v2-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| row_status | [RowStatus](#slash-api-v2-RowStatus) |  |  |
| created_ts | [int64](#int64) |  |  |
| updated_ts | [int64](#int64) |  |  |
| role | [Role](#slash-api-v2-Role) |  |  |
| email | [string](#string) |  |  |
| nickname | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="slash-api-v2-UserAccessToken"></a>

### UserAccessToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |
| description | [string](#string) |  |  |
| issued_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| expires_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 


<a name="slash-api-v2-Role"></a>

### Role


| Name | Number | Description |
| ---- | ------ | ----------- |
| ROLE_UNSPECIFIED | 0 |  |
| ADMIN | 1 |  |
| USER | 2 |  |


 

 


<a name="slash-api-v2-UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListUsers | [ListUsersRequest](#slash-api-v2-ListUsersRequest) | [ListUsersResponse](#slash-api-v2-ListUsersResponse) | ListUsers returns a list of users. |
| GetUser | [GetUserRequest](#slash-api-v2-GetUserRequest) | [GetUserResponse](#slash-api-v2-GetUserResponse) | GetUser returns a user by id. |
| CreateUser | [CreateUserRequest](#slash-api-v2-CreateUserRequest) | [CreateUserResponse](#slash-api-v2-CreateUserResponse) | CreateUser creates a new user. |
| DeleteUser | [DeleteUserRequest](#slash-api-v2-DeleteUserRequest) | [DeleteUserResponse](#slash-api-v2-DeleteUserResponse) | DeleteUser deletes a user by id. |
| ListUserAccessTokens | [ListUserAccessTokensRequest](#slash-api-v2-ListUserAccessTokensRequest) | [ListUserAccessTokensResponse](#slash-api-v2-ListUserAccessTokensResponse) | ListUserAccessTokens returns a list of access tokens for a user. |
| CreateUserAccessToken | [CreateUserAccessTokenRequest](#slash-api-v2-CreateUserAccessTokenRequest) | [CreateUserAccessTokenResponse](#slash-api-v2-CreateUserAccessTokenResponse) | CreateUserAccessToken creates a new access token for a user. |
| DeleteUserAccessToken | [DeleteUserAccessTokenRequest](#slash-api-v2-DeleteUserAccessTokenRequest) | [DeleteUserAccessTokenResponse](#slash-api-v2-DeleteUserAccessTokenResponse) | DeleteUserAccessToken deletes an access token for a user. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

