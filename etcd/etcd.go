package etcd

import (
	"crypto/tls"
	"crypto/x509"
)

// NewTLSConfig 构造一个https的配置对象
func NewTLSConfig(caCert, cert, privateKey []byte) (*tls.Config, error) {
	tlsCert, err := tls.X509KeyPair(cert, privateKey)
	if err != nil {
		return nil, err
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		RootCAs:      pool,
	}
	return tlsConfig, nil
}
