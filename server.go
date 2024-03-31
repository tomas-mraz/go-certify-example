package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/johanbrandhorst/certify"
	"github.com/johanbrandhorst/certify/issuers/vault"
	"log"
	"net"
	"net/http"
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

	c := &certify.Certify{
		CommonName:  "uw.com.uk",
		Issuer:      issuer,
		Cache:       certify.NewMemCache(),
		RenewBefore: 24 * time.Hour,
		CertConfig: &certify.CertConfig{
			SubjectAlternativeNames:   []string{"uw.com"},
			IPSubjectAlternativeNames: []net.IP{net.IPv6loopback},
		},
	}

	certpool := x509.NewCertPool()
	if !certpool.AppendCertsFromPEM(pem) {
		log.Fatalf("Can't parse client certificate authority")
	}

	s := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			ClientAuth:           tls.RequireAndVerifyClientCert,
			ClientCAs:            certpool,
			GetCertificate:       c.GetCertificate,
			GetClientCertificate: c.GetClientCertificate,
		},
	}

}
