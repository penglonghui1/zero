package soap

import (
	"bytes"
	"encoding/xml"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type SOAPEnvelope1 struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Header  *SOAPHeader1
	Body    SOAPBody1
}

type SOAPHeader1 struct {
	XMLName xml.Name      `xml:"soap:Header"`
	Soap    string        `xml:"xmlns:soap,attr,omitempty"`
	Xsd     string        `xml:"xsd,attr,omitempty"`
	Xsi     string        `xml:"xsi,attr,omitempty"`
	Fault   *SOAPFault1   `xml:",omitempty"`
	Content interface{}   `xml:",omitempty"`
	Items   []interface{} `xml:",omitempty"`
}

type SOAPBody1 struct {
	XMLName xml.Name    `xml:"soap:Body"`
	Soap    string      `xml:"xmlns:soap,attr,omitempty"`
	Xsd     string      `xml:"xsd,attr,omitempty"`
	Xsi     string      `xml:"xsi,attr,omitempty"`
	Fault   *SOAPFault1 `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault1 struct {
	XMLName xml.Name `xml:"Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

func (b *SOAPBody1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault1{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault1) Error() string {
	return f.String
}

func (s *SOAPClient) CallPatch(soapAction string, request SOAPEnvelope1, response interface{}) error {
	if s.headers != nil && len(s.headers) > 0 {
		request.Header = &SOAPHeader1{Items: s.headers}
	}

	bytess, err := xml.Marshal(request)
	if err != nil {
		return err
	}

	byteStr := strings.ReplaceAll(string(bytess), "&gt;", ">")

	buffer := new(bytes.Buffer)
	buffer.WriteString(byteStr)

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: s.tlsCfg,
		Dial:            dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	rawstr := string(rawbody)
	rawstr = html.UnescapeString(rawstr)
	log.Println("response raw -- ", rawstr)

	if err := (SOAPEnvelope{}).Scan([]byte(rawstr), response); err != nil {
		return err
	}

	return nil
}

// Scan 扫描参数到结构体
func (s SOAPEnvelope1) Scan(rawBody []byte, dest interface{}) error {
	s.Body.Content = dest
	return xml.Unmarshal(rawBody, &s)
}
