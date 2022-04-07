package client

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/coreos/etcd/clientv3"
)

type EtcdConfig struct {
}

type EtcdProxy struct {
	etcdCli clientv3.Client
}

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
