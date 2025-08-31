package client

import (
	credential "github.com/rancher/muchang/credentials"
	models "github.com/rancher/muchang/darabonba-openapi/models"
	openapiutil "github.com/rancher/muchang/darabonba-openapi/utils"
	spi "github.com/rancher/muchang/gateway-spi/client"
	"github.com/rancher/muchang/utils/tea/dara"
)

type Config = models.Config
type GlobalParameters = models.GlobalParameters
type Params = models.Params
type OpenApiRequest = models.OpenApiRequest
type Client struct {
	DisableSDKError      *bool
	Endpoint             *string
	RegionId             *string
	Protocol             *string
	Method               *string
	UserAgent            *string
	EndpointRule         *string
	EndpointMap          map[string]*string
	Suffix               *string
	ReadTimeout          *int
	ConnectTimeout       *int
	HttpProxy            *string
	HttpsProxy           *string
	Socks5Proxy          *string
	Socks5NetWork        *string
	NoProxy              *string
	Network              *string
	ProductId            *string
	MaxIdleConns         *int
	EndpointType         *string
	OpenPlatformEndpoint *string
	Credential           credential.Credential
	SignatureVersion     *string
	SignatureAlgorithm   *string
	Headers              map[string]*string
	Spi                  spi.ClientInterface
	GlobalParameters     *openapiutil.GlobalParameters
	Key                  *string
	Cert                 *string
	Ca                   *string
	DisableHttp2         *bool
	RetryOptions         *dara.RetryOptions
	HttpClient           dara.HttpClient
	TlsMinVersion        *string
	AttributeMap         *spi.AttributeMap
}

// Description:
//
// # Init client with Config
//
// @param config - config contains the necessary information to create a client
func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	if dara.IsNil(config) {
		_err = &ClientError{
			Code:    dara.String("ParameterMissing"),
			Message: dara.String("'config' can not be unset"),
		}
		return _err
	}

	if (!dara.IsNil(config.AccessKeyId) && dara.StringValue(config.AccessKeyId) != "") && (!dara.IsNil(config.AccessKeySecret) && dara.StringValue(config.AccessKeySecret) != "") {
		if !dara.IsNil(config.SecurityToken) && dara.StringValue(config.SecurityToken) != "" {
			config.Type = dara.String("sts")
		} else {
			config.Type = dara.String("access_key")
		}

		credentialConfig := &credential.Config{
			AccessKeyId:     config.AccessKeyId,
			Type:            config.Type,
			AccessKeySecret: config.AccessKeySecret,
		}
		credentialConfig.SecurityToken = config.SecurityToken
		client.Credential, _err = credential.NewCredential(credentialConfig)
		if _err != nil {
			return _err
		}

	} else if !dara.IsNil(config.BearerToken) && dara.StringValue(config.BearerToken) != "" {
		cc := &credential.Config{
			Type:        dara.String("bearer"),
			BearerToken: config.BearerToken,
		}
		client.Credential, _err = credential.NewCredential(cc)
		if _err != nil {
			return _err
		}

	} else if !dara.IsNil(config.Credential) {
		client.Credential = config.Credential
	}

	client.Endpoint = config.Endpoint
	client.EndpointType = config.EndpointType
	client.Network = config.Network
	client.Suffix = config.Suffix
	client.Protocol = config.Protocol
	client.Method = config.Method
	client.RegionId = config.RegionId
	client.UserAgent = config.UserAgent
	client.ReadTimeout = config.ReadTimeout
	client.ConnectTimeout = config.ConnectTimeout
	client.HttpProxy = config.HttpProxy
	client.HttpsProxy = config.HttpsProxy
	client.NoProxy = config.NoProxy
	client.Socks5Proxy = config.Socks5Proxy
	client.Socks5NetWork = config.Socks5NetWork
	client.MaxIdleConns = config.MaxIdleConns
	client.SignatureVersion = config.SignatureVersion
	client.SignatureAlgorithm = config.SignatureAlgorithm
	client.GlobalParameters = config.GlobalParameters
	client.Key = config.Key
	client.Cert = config.Cert
	client.Ca = config.Ca
	client.DisableHttp2 = config.DisableHttp2
	client.RetryOptions = config.RetryOptions
	client.HttpClient = config.HttpClient
	client.TlsMinVersion = config.TlsMinVersion
	return nil
}

func (client *Client) GetAccessDeniedDetail(err map[string]interface{}) (_result map[string]interface{}) {
	var accessDeniedDetail map[string]interface{}
	if !dara.IsNil(err["AccessDeniedDetail"]) {
		detail1 := dara.ToMap(err["AccessDeniedDetail"])
		accessDeniedDetail = detail1
	} else if !dara.IsNil(err["accessDeniedDetail"]) {
		detail2 := dara.ToMap(err["accessDeniedDetail"])
		accessDeniedDetail = detail2
	}

	_result = accessDeniedDetail
	return _result
}

// Description:
//
// get RPC header for debug
func (client *Client) GetRpcHeaders() (_result map[string]*string, _err error) {
	headers := client.Headers
	client.Headers = nil
	_result = headers
	return _result, _err
}

// Description:
//
// # If the endpointRule and config.endpoint are empty, throw error
//
// @param config - config contains the necessary information to create a client
func (client *Client) CheckConfig(config *openapiutil.Config) (_err error) {
	if dara.IsNil(client.EndpointRule) && dara.IsNil(config.Endpoint) {
		_err = &ClientError{
			Code:    dara.String("ParameterMissing"),
			Message: dara.String("'config.endpoint' can not be empty"),
		}
		return _err
	}

	return _err
}
