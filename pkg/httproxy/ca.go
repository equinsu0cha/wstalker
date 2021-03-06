/*
Released as open source by NCC Group Plc - http://www.nccgroup.com/

Developed by Jose Selvi, jose dot selvi at nccgroup dot com

http://www.github.com/nccgroup/wstalker

Released under AGPL see LICENSE for more information
*/

package httproxy

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/elazarl/goproxy"
)

var caCert = []byte(`-----BEGIN CERTIFICATE-----
MIIEvDCCAqQCCQCD9c6+9fhpzzANBgkqhkiG9w0BAQsFADAgMQswCQYDVQQGEwJV
SzERMA8GA1UEAwwId3N0YWxrZXIwHhcNMTkwNjA3MTcyMzA3WhcNMTkwNzA3MTcy
MzA3WjAgMQswCQYDVQQGEwJVSzERMA8GA1UEAwwId3N0YWxrZXIwggIiMA0GCSqG
SIb3DQEBAQUAA4ICDwAwggIKAoICAQDFJA9cRVtsKCCFTPXAQTtPpfPACxddbRQf
wVwBMkwKjkobtVwfD8OG+Mr5rrDRikuaSBTOHof7Se6SQwpx4zWMM9xRWzvRaUeo
Dk6tJE7+pGQt+O5quqM+UFIGmlq2s4E/1eDIiaAHCFfa8CnZ5DbpHKug5CO+C3NJ
XV3A71CZjgMbel4xZd09BR6/yfGf8SB35R70aaJRpNUMdKSUU9PxzsXVFj/2EQLs
4PCycBDRRfJgLfitjrUUM6G7Qh42rKSW6oEFdBA1T6J3lr6jR03Ir9QxxAVbAfj3
43tGflYpJOQKL+ilJn+0U/H+yGTtvoDmZSkhK/IY/TEUrV8B3TuzYJm1ACkKArFI
MzplZ+YRCnOp/nE+Ohf2pAEV8ZEmlVuuGy3JN+GxvkHUwFtXXyXPtZ0eX4O8RtlF
oK7ZntoBRHwbBYtIJg7C1jWLLc9BhYAeAO6vT3kup1UCormvKaaxUuELkztaLqEf
8GXoQTyBVSwaDGHuJv+liFambNh20tm068844mjbXnLqRT6BGjlhSIz9/v3ElGfV
c22Ewer3xUZnmDUqzaljgQX5K+OGl7xW1Yrk0AFl/Enh7k0D8h9aSwEusbsUqvY+
8kTP7kN+qmuDSSGFsnJp0j1Zq7nzxpRelaUC2mFfvBqZWFkocwjmWLwFb6zWOb/J
1BP3C0OzlQIDAQABMA0GCSqGSIb3DQEBCwUAA4ICAQCHjpCxwwuY0NHS4FoMhNCd
gwfL6rcp4E2l+IpuaeYraqEA/suofxlh5LAWcqRqZH7S+XnOU5z+zl7thZp7lDJ1
uVVDpv0TqEUbveutn9kpJ7/DPF/AJ+BJMM6smoRH3WxYxJ8Zv129OoYtNh8SgkZq
tsp8NK3GbGqX7xbwBwmzrBvwwitIYuC28RB4s3/qVyIMfJ/nUN81HPROOZRL6UGi
TN2C3PLUtdHvLFZY9TzrUvK5aS8+oFAasaa79XcW6AFscARhHw2RdquWFl7tSL2v
zTD2dgtFzbwCx22dDzPWVLB8K4neuB6Kvm2ChivN+MrzUiplZbRGHQvIED0I7H8Y
3T3raTfuLSgLFW4x4Eq1J3ZjiCWV/hwWqistGwR6eWCmkFumc4Cv1SDdFDdpbT6B
xDyKqvJpsrjjf/tyU1qDBlGZ2VKYQ/Dh7VmtCTWE2CwkWsywOSApjlPdzVgQqQJh
LeQhVyzih1qJ24oBQP9xPcJM2aXzDvaP25zyrCo/Bz4Y1n9evQ0hhVH0r29LOEQC
4GOVk5gCvDyi8IVBr5Y//5pLq4dPlAo4kKdWmSjCB4zhawkhH3st6Z7VsKzWelAR
TjHiABHR5Y27cpPgAFe1bUoW25dpKqpaXo5WvuP3N7LSmRckIfQsv/EoqWTIK+dM
a1HDBNJoeWpbflmbuJZqCg==
-----END CERTIFICATE-----`)

// Intentionally left here. Check README.md for more information.
var caKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)

func (h *HttProxy) setCA(caCert, caKey []byte) error {
	goproxyCa, err := tls.X509KeyPair(caCert, caKey)
	if err != nil {
		return err
	}
	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
		return err
	}
	goproxy.GoproxyCa = goproxyCa
	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	return nil
}
