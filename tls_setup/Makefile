.PHONY: cfssl ca req clean
  
all: ca req

clean:
	rm -rf certs

preq:
	go get -u -tags nopkcs11 github.com/cloudflare/cfssl/cmd/cfssl
	go get -u github.com/cloudflare/cfssl/cmd/cfssljson

ca:
	mkdir -p certs
	cfssl gencert -initca config/ca_csr.json | cfssljson -bare certs/ca
	mv certs/ca-key.pem certs/ca.key

req:
	cfssl gencert \
	  -ca certs/ca.pem \
	  -ca-key certs/ca.key \
	  -config config/ca_cfg.json config/req_csr.json | cfssljson -bare certs/etcd1
	mv certs/etcd1-key.pem certs/etcd1.key


	cfssl gencert \
	  -ca certs/ca.pem \
	  -ca-key certs/ca.key \
	  -config config/ca_cfg.json config/req_csr.json | cfssljson -bare certs/etcd2
	mv certs/etcd2-key.pem certs/etcd2.key

	cfssl gencert \
	  -ca certs/ca.pem \
	  -ca-key certs/ca.key \
	  -config config/ca_cfg.json config/req_csr.json | cfssljson -bare certs/etcd3
	mv certs/etcd3-key.pem certs/etcd3.key

.PHONY: clean preq ca req
