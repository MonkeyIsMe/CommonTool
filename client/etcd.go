package client

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// EtcdConfig etcd的配置文件
type EtcdConfig struct {
	endpoints        []string
	dialTimeout      int
	autoSyncInterval int
	caCert           []byte
	cert             []byte
	privateKey       []byte
}

// EtcdProxy etcd的proxy
type EtcdProxy struct {
	etcdCli clientv3.Client
	callee  string
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

// NewEtcdClient 初始化一个etcd的客户端
func NewEtcdClient(etcdConfig EtcdConfig, connType string) (*clientv3.Client, error) {
	const httpsValue = "https"
	var tlsConfig *tls.Config
	var err error

	if connType == httpsValue {
		tlsConfig, err = NewTLSConfig(etcdConfig.caCert, etcdConfig.cert, etcdConfig.privateKey)
		if err != nil {
			log.Fatalf("init https tls config err: [%+v]", err)
			return nil, err
		}
	}

	return clientv3.New(clientv3.Config{
		Endpoints:        etcdConfig.endpoints,
		DialTimeout:      time.Duration(etcdConfig.dialTimeout) * time.Millisecond,
		TLS:              tlsConfig,
		AutoSyncInterval: time.Duration(etcdConfig.autoSyncInterval) * time.Millisecond,
	})
}
