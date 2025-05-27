package wx

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/shenghui0779/yiigo"
)

// Client is the interface that do http request
type Client interface {
	// Post sends an HTTP post request
	Do(ctx context.Context, method, reqURL string, body []byte, options ...yiigo.HTTPOption) ([]byte, error)

	// Upload sends an HTTP post request for uploading media
	Upload(ctx context.Context, reqURL string, form yiigo.UploadForm, options ...yiigo.HTTPOption) ([]byte, error)

	// Set sets options for client
	Set(options ...ClientOption)
}

type wxclient struct {
	client yiigo.HTTPClient
	logger Logger
	debug  bool
}

func (c *wxclient) Do(ctx context.Context, method, reqURL string, body []byte, options ...yiigo.HTTPOption) ([]byte, error) {
	logData := &LogData{
		URL:    reqURL,
		Method: method,
		Body:   body,
	}

	now := time.Now().Local()

	defer func() {
		logData.Duration = time.Since(now)

		if c.debug {
			c.logger.Log(ctx, logData)
		}
	}()

	resp, err := c.client.Do(ctx, method, reqURL, body, options...)

	if err != nil {
		logData.Error = err

		return nil, err
	}

	defer resp.Body.Close()

	logData.StatusCode = resp.StatusCode

	if resp.StatusCode >= http.StatusBadRequest {
		io.Copy(ioutil.Discard, resp.Body)

		return nil, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logData.Error = err

		return nil, err
	}

	logData.Response = b

	return b, nil
}

func (c *wxclient) Upload(ctx context.Context, reqURL string, form yiigo.UploadForm, options ...yiigo.HTTPOption) ([]byte, error) {
	logData := &LogData{
		URL:    reqURL,
		Method: http.MethodPost,
	}

	now := time.Now().Local()

	defer func() {
		logData.Duration = time.Since(now)

		if c.debug {
			c.logger.Log(ctx, logData)
		}
	}()

	resp, err := c.client.Upload(ctx, reqURL, form, options...)

	if err != nil {
		logData.Error = err

		return nil, err
	}

	defer resp.Body.Close()

	logData.StatusCode = resp.StatusCode

	if resp.StatusCode >= http.StatusBadRequest {
		io.Copy(ioutil.Discard, resp.Body)

		return nil, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logData.Error = err

		return nil, err
	}

	logData.Response = b

	return b, nil
}

func (c *wxclient) Set(options ...ClientOption) {
	for _, f := range options {
		f(c)
	}
}

// DefaultClient returns a new default wechat client
func DefaultClient(certs ...tls.Certificate) Client {
	tlscfg := &tls.Config{
		InsecureSkipVerify: true,
	}

	if len(certs) != 0 {
		tlscfg.Certificates = certs
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 60 * time.Second,
			}).DialContext,
			TLSClientConfig:       tlscfg,
			MaxIdleConns:          0,
			MaxIdleConnsPerHost:   1000,
			MaxConnsPerHost:       1000,
			IdleConnTimeout:       60 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	return &wxclient{
		client: yiigo.NewHTTPClient(client),
		logger: DefaultLogger(),
	}
}

func NewProxyHTTPClient(proxyURL string, certs ...tls.Certificate) Client {
	// 1. 初始化 TLS 配置
	tlscfg := &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证（生产环境应谨慎使用）
	}

	if len(certs) > 0 {
		tlscfg.Certificates = certs
	}

	// 2. 解析代理URL（如果提供）

	var proxyFunc func(*http.Request) (*url.URL, error)
	proxyFunc = http.ProxyFromEnvironment // 默认使用环境变量代理

	if proxyURL != "" {
		parsedProxyURL, err := url.Parse(proxyURL)
		if err != nil {
			return nil
		}
		proxyFunc = http.ProxyURL(parsedProxyURL)
	}

	// 3. 创建自定义 Transport
	transport := &http.Transport{
		Proxy: proxyFunc,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		TLSClientConfig:       tlscfg,
		MaxIdleConns:          0,    // 0 表示无限制
		MaxIdleConnsPerHost:   1000, // 默认值
		MaxConnsPerHost:       1000, // 默认值
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// 4. 创建 HTTP Client
	client := &http.Client{
		Transport: transport,
	}

	// 5. 返回微信客户端
	return &wxclient{
		client: yiigo.NewHTTPClient(client),
		logger: DefaultLogger(),
	}
}
