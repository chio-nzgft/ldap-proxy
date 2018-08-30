package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	ldapc "github.com/mkorolyov/ldap-proxy/go-ldapc"
	ldap "github.com/mkorolyov/ldap-proxy/ldapserver"
)

type Config struct {
	ListenAddr           string
	AdserverIP           string
	Adserverport         string
	AdserverBindDN       string
	AdserverBaseDN       string
	AdserverBindPassword string
	AdserverFilter       string
}

var c Config

func main() {
	//read config
	r, err := os.Open("conf.json")
	if err != nil {
		log.Fatalln(err)
	}
	decoder := json.NewDecoder(r)

	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalln(err)
	}

	//ldap logger
	ldap.Logger = log.New(os.Stdout, "[server] ", log.LstdFlags)

	//Create a new LDAP Server
	server := ldap.NewServer()

	routes := ldap.NewRouteMux()
	routes.Bind(handleBind)
	server.Handle(routes)

	// listen on 10389
	go server.ListenAndServe(c.ListenAddr)

	// When CTRL+C, SIGINT and SIGTERM signal occurs
	// Then stop server gracefully
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	close(ch)

	server.Stop()
}

func handleBind(w ldap.ResponseWriter, m *ldap.Message) {
	var iPort int
	r := m.GetBindRequest()
	res := ldap.NewBindResponse(ldap.LDAPResultSuccess)
	username := string(r.Name())
	password := string(r.AuthenticationSimple())
	log.Printf("AD server %s:%s", c.AdserverIP, c.Adserverport)
	log.Printf("user %s", username)
	log.Printf("pass %s", password)
	iPort, _ = strconv.Atoi(c.Adserverport)

	ldapclient := &ldapc.Client{
		Protocol:  ldapc.LDAP,
		Host:      c.AdserverIP,
		Port:      iPort,
		TLSConfig: nil,
		Bind: &ldapc.AuthBind{
			BindDN:       c.AdserverBindDN,
			BindPassword: c.AdserverBindPassword,
			BaseDN:       c.AdserverBaseDN,
			Filter:       c.AdserverFilter,
		},
	}
	entry, err := ldapclient.Authenticate(username, password)
	if err != nil {
		log.Printf("LDAP Authenticate failed: %v\n", err)
		res.SetResultCode(ldap.LDAPResultInvalidCredentials)
		res.SetDiagnosticMessage("system error !!")
		w.Write(res)
		return
	} else {
		w.Write(res)
		log.Printf("%+v\n", entry)
		return
	}
	return
}
