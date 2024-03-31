package main

import (
	"context"
	"crypto/x509"
	"fmt"
	"github.com/johanbrandhorst/certify"
	"github.com/johanbrandhorst/certify/issuers/vault"
	"net"
	"net/url"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	issuer := &vault.Issuer{
		URL: &url.URL{
			Scheme: "https",
			Host:   "secure.cubyte.online",
		},
		//AuthMethod: &vault.AuthMethod(token),
		Token:      "token",
		Role:       "role",
		TimeToLive: time.Minute * 10,
		// Format is "<type_id>;utf8:<value>", where type_id
		// is an ASN.1 object identifier.
		//OtherSubjectAlternativeNames: []string{"1.3.6.1.4.1.311.20.2.3;utf8:devops@nope.com"},
		OtherSubjectAlternativeNames: nil,
		//URISubjectAlternativeNames:   []string{"spiffe://hostname/testing"},
	}

	conf := &certify.CertConfig{
		SubjectAlternativeNames:   []string{"uw.com"},
		IPSubjectAlternativeNames: []net.IP{net.IPv6loopback},
	}

	c := &certify.Certify{
		CommonName:  "uw.com.uk",
		Cache:       certify.NewMemCache(),
		RenewBefore: 24 * time.Hour,
		CertConfig:  conf,
	}

	s :=

	cn := "somename.com"

	tlsCert, err1 := iss.Issue(context.Background(), cn, conf)
	caCert, err2 := x509.ParseCertificate(tlsCert.Certificate[1])

}
