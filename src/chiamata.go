package main

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    // carico i certificati client per la mTLS
    cert, err := tls.LoadX509KeyPair("../cert/SVILUPPO_2018/SVILUPPO_2018_CERT.pem", "../cert/SVILUPPO_2018/SVILUPPO_2018.pem")
    if err != nil {
        panic(err)
    }

    // carico il certificato della CA
    caCert, err := ioutil.ReadFile("../cert/PopsoRootCA01.pem")
    if err != nil {
        panic(err)
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    config := &tls.Config{
        Certificates: []tls.Certificate{cert},
        RootCAs:      caCertPool,
        InsecureSkipVerify: false,
    }

    transport := &http.Transport{
        TLSClientConfig: config,
    }

    client := &http.Client{
        Transport: transport,
    }

    // per funzionare la chiamata deve passare gli header http corretti
    headers := http.Header{}
    headers.Set("Content-type", "text/xml")
    headers.Set("SoapAction", "http://scrittura.iuvonline.nodospcit.ws.popso.it/v1/getRTRead")

    // carico la busta soap dal file
    file, err := os.Open("../request_body.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // eseguo la POST
    req, err := http.NewRequest("POST", "POPSO_ENDPOINT", file)
    if err != nil {
        panic(err)
    }
    req.Header = headers

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    // stampo il risultato
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(body))
}
