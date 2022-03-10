package cwhydra

import "time"

type OAuth2Client struct {
	AllowedCorsOrigins []string `json:"allowed_cors_origins"`
	Audience           []string `json:"audience"`
	// Boolean value specifying whether the RP requires that a sid (session ID) Claim be included in the Logout Token to identify the RP session with the OP when the backchannel_logout_uri is used. If omitted, the default value is false.
	BackchannelLogoutSessionRequired bool `json:"backchannel_logout_session_required"`
	// RP URL that will cause the RP to log itself out when sent a Logout Token by the OP.
	BackchannelLogoutUri string `json:"backchannel_logout_uri"`
	// ID  is the id for this client.
	ClientId string `json:"client_id"`
	// Name is the human-readable string name of the client to be presented to the end-user during authorization.
	ClientName string `json:"client_name"`
	// Secret is the client's secret. The secret will be included in the create request as cleartext, and then never again. The secret is stored using BCrypt so it is impossible to recover it. Tell your users that they need to write the secret down as it will not be made available again.
	ClientSecret string `json:"client_secret"`
	// SecretExpiresAt is an integer holding the time at which the client secret will expire or 0 if it will not expire. The time is represented as the number of seconds from 1970-01-01T00:00:00Z as measured in UTC until the date/time of expiration.  This feature is currently not supported and it's value will always be set to 0.
	ClientSecretExpiresAt int64 `json:"client_secret_expires_at"`
	// ClientURI is an URL string of a web page providing information about the client. If present, the server SHOULD display this URL to the end-user in a clickable fashion.
	ClientUri string   `json:"client_uri"`
	Contacts  []string `json:"contacts"`
	// CreatedAt returns the timestamp of the client's creation.
	CreatedAt time.Time `json:"created_at"`
	// Boolean value specifying whether the RP requires that iss (issuer) and sid (session ID) query parameters be included to identify the RP session with the OP when the frontchannel_logout_uri is used. If omitted, the default value is false.
	FrontchannelLogoutSessionRequired bool `json:"frontchannel_logout_session_required"`
	// RP URL that will cause the RP to log itself out when rendered in an iframe by the OP. An iss (issuer) query parameter and a sid (session ID) query parameter MAY be included by the OP to enable the RP to validate the request and to determine which of the potentially multiple sessions is to be logged out; if either is included, both MUST be.
	FrontchannelLogoutUri string                 `json:"frontchannel_logout_uri"`
	GrantTypes            []string               `json:"grant_types"`
	Jwks                  map[string]interface{} `json:"jwks"`
	// URL for the Client's JSON Web Key Set [JWK] document. If the Client signs requests to the Server, it contains the signing key(s) the Server uses to validate signatures from the Client. The JWK Set MAY also contain the Client's encryption keys(s), which are used by the Server to encrypt responses to the Client. When both signing and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced JWK Set to indicate each key's intended usage. Although some algorithms allow the same key to be used for both signatures and encryption, doing so is NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used to provide X.509 representations of keys provided. When used, the bare key values MUST still be present and MUST match those in the certificate.
	JwksUri string `json:"jwks_uri"`
	// LogoURI is an URL string that references a logo for the client.
	LogoUri  string                 `json:"logo_uri"`
	Metadata map[string]interface{} `json:"metadata"`
	// Owner is a string identifying the owner of the OAuth 2.0 Client.
	Owner string `json:"owner"`
	// PolicyURI is a URL string that points to a human-readable privacy policy document that describes how the deployment organization collects, uses, retains, and discloses personal data.
	PolicyUri              string   `json:"policy_uri"`
	PostLogoutRedirectUris []string `json:"post_logout_redirect_uris"`
	RedirectUris           []string `json:"redirect_uris"`
	// RegistrationAccessToken can be used to update, get, or delete the OAuth2 Client.
	RegistrationAccessToken string `json:"registration_access_token"`
	// RegistrationClientURI is the URL used to update, get, or delete the OAuth2 Client.
	RegistrationClientUri string `json:"registration_client_uri"`
	// JWS [JWS] alg algorithm [JWA] that MUST be used for signing Request Objects sent to the OP. All Request Objects from this Client MUST be rejected, if not signed with this algorithm.
	RequestObjectSigningAlg string   `json:"request_object_signing_alg"`
	RequestUris             []string `json:"request_uris"`
	ResponseTypes           []string `json:"response_types"`
	// Scope is a string containing a space-separated list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client can use when requesting access tokens.
	Scope string `json:"scope"`
	// URL using the https scheme to be used in calculating Pseudonymous Identifiers by the OP. The URL references a file with a single JSON array of redirect_uri values.
	SectorIdentifierUri string `json:"sector_identifier_uri"`
	// SubjectType requested for responses to this Client. The subject_types_supported Discovery parameter contains a list of the supported subject_type values for this server. Valid types include `pairwise` and `public`.
	SubjectType string `json:"subject_type"`
	// Requested Client Authentication method for the Token Endpoint. The options are client_secret_post, client_secret_basic, private_key_jwt, and none.
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method"`
	// Requested Client Authentication signing algorithm for the Token Endpoint.
	TokenEndpointAuthSigningAlg string `json:"token_endpoint_auth_signing_alg"`
	// TermsOfServiceURI is a URL string that points to a human-readable terms of service document for the client that describes a contractual relationship between the end-user and the client that the end-user accepts when authorizing the client.
	TosUri string `json:"tos_uri"`
	// UpdatedAt returns the timestamp of the last update.
	UpdatedAt time.Time `json:"updated_at"`
	// JWS alg algorithm [JWA] REQUIRED for signing UserInfo Responses. If this is specified, the response will be JWT [JWT] serialized, and signed using JWS. The default, if omitted, is for the UserInfo Response to return the Claims as a UTF-8 encoded JSON object using the application/json content-type.
	UserinfoSignedResponseAlg string `json:"userinfo_signed_response_alg"`
}

type OpenIDConnectContext struct {
	// ACRValues is the Authentication AuthorizationContext Class Reference requested in the OAuth 2.0 Authorization request. It is a parameter defined by OpenID Connect and expresses which level of authentication (e.g. 2FA) is required.  OpenID Connect defines it as follows: > Requested Authentication AuthorizationContext Class Reference values. Space-separated string that specifies the acr values that the Authorization Server is being requested to use for processing this Authentication Request, with the values appearing in order of preference. The Authentication AuthorizationContext Class satisfied by the authentication performed is returned as the acr Claim Value, as specified in Section 2. The acr Claim is requested as a Voluntary Claim by this parameter.
	AcrValues []string `json:"acr_values"`
	// Display is a string value that specifies how the Authorization Server displays the authentication and consent user interface pages to the End-User. The defined values are: page: The Authorization Server SHOULD display the authentication and consent UI consistent with a full User Agent page view. If the display parameter is not specified, this is the default display mode. popup: The Authorization Server SHOULD display the authentication and consent UI consistent with a popup User Agent window. The popup User Agent window should be of an appropriate size for a login-focused dialog and should not obscure the entire window that it is popping up over. touch: The Authorization Server SHOULD display the authentication and consent UI consistent with a device that leverages a touch interface. wap: The Authorization Server SHOULD display the authentication and consent UI consistent with a \"feature phone\" type display.  The Authorization Server MAY also attempt to detect the capabilities of the User Agent and present an appropriate display.
	Display string `json:"display"`
	// IDTokenHintClaims are the claims of the ID Token previously issued by the Authorization Server being passed as a hint about the End-User's current or past authenticated session with the Client.
	IdTokenHintClaims map[string]map[string]interface{} `json:"id_token_hint_claims"`
	// LoginHint hints about the login identifier the End-User might use to log in (if necessary). This hint can be used by an RP if it first asks the End-User for their e-mail address (or other identifier) and then wants to pass that value as a hint to the discovered authorization service. This value MAY also be a phone number in the format specified for the phone_number Claim. The use of this parameter is optional.
	LoginHint string `json:"login_hint"`
	// UILocales is the End-User'id preferred languages and scripts for the user interface, represented as a space-separated list of BCP47 [RFC5646] language tag values, ordered by preference. For instance, the value \"fr-CA fr en\" represents a preference for French as spoken in Canada, then French (without a region designation), followed by English (without a region designation). An error SHOULD NOT result if some or all of the requested locales are not supported by the OpenID Provider.
	UiLocales []string `json:"ui_locales"`
}

/* REJECT */

type RejectRequest struct {
	// The error should follow the OAuth2 error format (e.g. `invalid_request`, `login_required`).  Defaults to `request_denied`.
	Error string `json:"error"`
	// Debug contains information to help resolve the problem as a developer. Usually not exposed to the public but only in the server logs.
	ErrorDebug string `json:"error_debug"`
	// Description of the error in a human readable format.
	ErrorDescription string `json:"error_description"`
	// Hint to help resolve the error.
	ErrorHint string `json:"error_hint"`
	// Represents the HTTP status code of the error (e.g. 401 or 403)  Defaults to 400
	StatusCode int64 `json:"status_code"`
}

/* CONSENT REQUEST */

type ConsentRequest struct {
	// ACR represents the Authentication AuthorizationContext Class Reference value for this authentication session. You can use it to express that, for example, a user authenticated using two factor authentication.
	Acr string   `json:"acr"`
	Amr []string `json:"amr"`
	// ID is the identifier (\"authorization challenge\") of the consent authorization request. It is used to identify the session.
	Challenge string                 `json:"challenge"`
	Client    OAuth2Client           `json:"client"`
	Context   map[string]interface{} `json:"context"`
	// LoginChallenge is the login challenge this consent challenge belongs to. It can be used to associate a login and consent request in the login & consent app.
	LoginChallenge string `json:"login_challenge"`
	// LoginSessionID is the login session ID. If the user-agent reuses a login session (via cookie / remember flag) this ID will remain the same. If the user-agent did not have an existing authentication session (e.g. remember is false) this will be a new random value. This value is used as the \"sid\" parameter in the ID Token and in OIDC Front-/Back- channel logout. It's value can generally be used to associate consecutive login requests by a certain user.
	LoginSessionId string               `json:"login_session_id"`
	OidcContext    OpenIDConnectContext `json:"oidc_context"`
	// RequestURL is the original OAuth 2.0 Authorization URL requested by the OAuth 2.0 client. It is the URL which initiates the OAuth 2.0 Authorization Code or OAuth 2.0 Implicit flow. This URL is typically not needed, but might come in handy if you want to deal with additional request parameters.
	RequestUrl                   string   `json:"request_url"`
	RequestedAccessTokenAudience []string `json:"requested_access_token_audience"`
	RequestedScope               []string `json:"requested_scope"`
	// Skip, if true, implies that the client has requested the same scopes from the same user previously. If true, you must not ask the user to grant the requested scopes. You must however either allow or deny the consent request using the usual API call.
	Skip bool `json:"skip"`
	// Subject is the user ID of the end-user that authenticated. Now, that end user needs to grant or deny the scope requested by the OAuth 2.0 client.
	Subject string `json:"subject"`
}

type ConsentRequestSession struct {
	// AccessToken sets session data for the access and refresh token, as well as any future tokens issued by the refresh grant. Keep in mind that this data will be available to anyone performing OAuth 2.0 Challenge Introspection. If only your services can perform OAuth 2.0 Challenge Introspection, this is usually fine. But if third parties can access that endpoint as well, sensitive data from the session might be exposed to them. Use with care!
	AccessToken map[string]interface{} `json:"access_token"`
	// IDToken sets session data for the OpenID Connect ID token. Keep in mind that the session'id payloads are readable by anyone that has access to the ID Challenge. Use with care!
	IdToken map[string]interface{} `json:"id_token"`
}

type AcceptConsentRequest struct {
	GrantAccessTokenAudience []string  `json:"grant_access_token_audience"`
	GrantScope               []string  `json:"grant_scope"`
	HandledAt                time.Time `json:"handled_at"`
	// Remember, if set to true, tells ORY Hydra to remember this consent authorization and reuse it if the same client asks the same user for the same, or a subset of, scope.
	Remember bool `json:"remember"`
	// RememberFor sets how long the consent authorization should be remembered for in seconds. If set to `0`, the authorization will be remembered indefinitely.
	RememberFor int64                 `json:"remember_for"`
	Session     ConsentRequestSession `json:"session"`
}

/* LOGIN REQUEST */

type LoginRequest struct {
	// ID is the identifier (\"login challenge\") of the login request. It is used to identify the session.
	Challenge   string               `json:"challenge"`
	Client      OAuth2Client         `json:"client"`
	OidcContext OpenIDConnectContext `json:"oidc_context"`
	// RequestURL is the original OAuth 2.0 Authorization URL requested by the OAuth 2.0 client. It is the URL which initiates the OAuth 2.0 Authorization Code or OAuth 2.0 Implicit flow. This URL is typically not needed, but might come in handy if you want to deal with additional request parameters.
	RequestUrl                   string   `json:"request_url"`
	RequestedAccessTokenAudience []string `json:"requested_access_token_audience"`
	RequestedScope               []string `json:"requested_scope"`
	// SessionID is the login session ID. If the user-agent reuses a login session (via cookie / remember flag) this ID will remain the same. If the user-agent did not have an existing authentication session (e.g. remember is false) this will be a new random value. This value is used as the \"sid\" parameter in the ID Token and in OIDC Front-/Back- channel logout. It's value can generally be used to associate consecutive login requests by a certain user.
	SessionId string `json:"session_id"`
	// Skip, if true, implies that the client has requested the same scopes from the same user previously. If true, you can skip asking the user to grant the requested scopes, and simply forward the user to the redirect URL.  This feature allows you to update / set session information.
	Skip bool `json:"skip"`
	// Subject is the user ID of the end-user that authenticated. Now, that end user needs to grant or deny the scope requested by the OAuth 2.0 client. If this value is set and `skip` is true, you MUST include this subject type when accepting the login request, or the request will fail.
	Subject string `json:"subject"`
}

type AcceptLoginRequest struct {
	// ACR sets the Authentication AuthorizationContext Class Reference value for this authentication session. You can use it to express that, for example, a user authenticated using two factor authentication.
	Acr     string                 `json:"acr"`
	Amr     []string               `json:"amr"`
	Context map[string]interface{} `json:"context"`
	// ForceSubjectIdentifier forces the \"pairwise\" user ID of the end-user that authenticated. The \"pairwise\" user ID refers to the (Pairwise Identifier Algorithm)[http://openid.net/specs/openid-connect-core-1_0.html#PairwiseAlg] of the OpenID Connect specification. It allows you to set an obfuscated subject (\"user\") identifier that is unique to the client.  Please note that this changes the user ID on endpoint /userinfo and sub claim of the ID Token. It does not change the sub claim in the OAuth 2.0 Introspection.  Per default, ORY Hydra handles this value with its own algorithm. In case you want to set this yourself you can use this field. Please note that setting this field has no effect if `pairwise` is not configured in ORY Hydra or the OAuth 2.0 Client does not expect a pairwise identifier (set via `subject_type` key in the client's configuration).  Please also be aware that ORY Hydra is unable to properly compute this value during authentication. This implies that you have to compute this value on every authentication process (probably depending on the client ID or some other unique value).  If you fail to compute the proper value, then authentication processes which have id_token_hint set might fail.
	ForceSubjectIdentifier string `json:"force_subject_identifier"`
	// Remember, if set to true, tells ORY Hydra to remember this user by telling the user agent (browser) to store a cookie with authentication data. If the same user performs another OAuth 2.0 Authorization Request, he/she will not be asked to log in again.
	Remember bool `json:"remember"`
	// RememberFor sets how long the authentication should be remembered for in seconds. If set to `0`, the authorization will be remembered for the duration of the browser session (using a session cookie).
	RememberFor int64 `json:"remember_for"`
	// Subject is the user ID of the end-user that authenticated.
	Subject string `json:"subject"`
}

/* LOGOUT REQUEST */

type LogoutResponse struct {
	// Challenge is the identifier (\"logout challenge\") of the logout authentication request. It is used to identify the session.
	Challenge string       `json:"challenge"`
	Client    OAuth2Client `json:"client"`
	// RequestURL is the original Logout URL requested.
	RequestUrl string `json:"request_url"`
	// RPInitiated is set to true if the request was initiated by a Relying Party (RP), also known as an OAuth 2.0 Client.
	RpInitiated bool `json:"rp_initiated"`
	// SessionID is the login session ID that was requested to log out.
	Sid string `json:"sid"`
	// Subject is the user for whom the logout was request.
	Subject string `json:"subject"`
}
