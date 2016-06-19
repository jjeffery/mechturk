// Package soap provides SOAP support.
package soap

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

// Envelope is for serializing SOAP envelopes.
type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body
}

// Body is for serializing the body of SOAP envelopes.
type Body struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault   *Fault      `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

// Fault is for serializing SOAP faults.
type Fault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

// Client is a SOAP client.
type Client struct {
	Logger *log.Logger
	url    string
	tls    bool
}

func (b *Body) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				b.Fault = &Fault{}
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

func (f *Fault) Error() string {
	return f.String
}

func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}

func (s *Client) Call(soapAction string, request, response interface{}) error {
	envelope := Envelope{}
	envelope.Body.Content = request
	buffer := new(bytes.Buffer)
	encoder := xml.NewEncoder(buffer)
	if s.Logger != nil {
		encoder.Indent("    ", "    ")
	}

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	if s.Logger != nil {
		s.Logger.Println("SOAP request:", "\n", buffer.String())
	}

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "go-mturk/0.1")
	req.Close = true

	client := &http.Client{}
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
		if s.Logger != nil {
			s.Logger.Println("SOAP response: empty")
		}
		return nil
	}

	respEnvelope := new(Envelope)
	respEnvelope.Body = Body{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		if s.Logger != nil {
			s.Logger.Println("SOAP response:", err, ":\n", string(rawbody))
		}
		return err
	}

	if s.Logger != nil {
		s.Logger.Printf("SOAP response, HTTP status=%s:\n%s",
			res.Status,
			string(rawbody))
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
