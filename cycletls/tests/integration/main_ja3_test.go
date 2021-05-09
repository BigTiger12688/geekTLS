// +build integration

package cycletls_test

import (
	//"fmt"
	"encoding/json"
	"log"
	"testing"

	cycletls "github.com/Danny-Dasilva/CycleTLS/cycletls"
)

type CycleTLSOptions struct {
	Ja3Hash string `json:"ja3_hash"`
	Ja3 string `json:"ja3"`
	UserAgent string `json:"User-Agent"`
	HTTPResponse int

}

type Ja3erResp struct {
	Ja3Hash string `json:"ja3_hash"`
	Ja3 string `json:"ja3"`
	UserAgent string `json:"User-Agent"`

}

var CycleTLSResults = []CycleTLSOptions {
	{"aa7744226c695c0b2e440419848cf700", //
	"771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0", 
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", 
	200},
	{"bc6c386f480ee97b9d9e52d472b772d8", // HelloChrome_58
	"771,49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53-10,65281-0-23-35-13-5-18-16-11-10,29-23-24,0", 
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", 
	200},
	{"bc6c386f480ee97b9d9e52d472b772d8", // HelloChrome_62
	"771,49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53-10,65281-0-23-35-13-5-18-16-11-10,29-23-24,0", 
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", 
	200},
	{"b4918ee98d0f0deb4e48563ca749ef10", // HelloChrome_70
	"771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53-10,65281-0-23-35-13-5-18-16-11-51-45-43-10-27-21,29-23-24,0", 
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", 
	200},
	{"66918128f1b9b03303d77c6f2eefd128", // HelloChrome_72
	"771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53-10,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0", 
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", 
	200},
	{"b32309a26951912be7dba376398abc3b", // HelloChrome_83
	"771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0", 
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36", 
	200},
	{"0ffee3ba8e615ad22535e7f771690a28", // HelloFirefox_55
	"771,49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-13,29-23-24-25,0", 
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", 
	200},
	{"0ffee3ba8e615ad22535e7f771690a28", // HelloFirefox_56
	"771,49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-13,29-23-24-25,0", 
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", 
	200},
	{"b20b44b18b853ef29ab773e921b03422", // HelloFirefox_63
	"771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0", 
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", 
	200},
	{"b20b44b18b853ef29ab773e921b03422", // HelloFirefox_65
	"771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0", 
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", 
	200},
	{"a69708a64f853c3bcc214c2c5faf84f3", // HelloIOS_11_1
	"771,49196-49195-49188-49187-49162-49161-52393-49200-49199-49192-49191-49172-49171-52392-157-156-61-60-53-47,65281-0-23-13-5-13172-18-16-11-10,29-23-24-25,0", 
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", 
	200},
	{"5c118da645babe52f060d0754256a73c", // HelloIOS_12_1
	"771,49196-49195-49188-49187-49162-49161-52393-49200-49199-49192-49191-49172-49171-52392-157-156-61-60-53-47-49160-49170-10,65281-0-23-13-5-13172-18-16-11-10,29-23-24-25,0", 
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0", 
	200},	

}

// {"ja3_hash":"aa7744226c695c0b2e440419848cf700", "ja3": "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0", "User-Agent": "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0"}
func TestIntegration(t *testing.T) {
	client := cycletls.Init()
	for _, options := range CycleTLSResults {

		response, err := client.Do("https://ja3er.com/json", cycletls.Options{
			Ja3: options.Ja3,
			UserAgent: options.UserAgent,
				
		}, "GET");
		if err != nil {
			t.Fatal("Unmarshal Error")
		}

		if (response.Response.Status != options.HTTPResponse) {
			t.Fatal("Expected Result Not given")
		} else {
			log.Println("ja3er: ", response.Response.Status)
		}
		ja3resp := new(Ja3erResp)

		err = json.Unmarshal([]byte(response.Response.Body), &ja3resp)
		if err != nil {
			t.Fatal("Unmarshal Error")
		}
		

		if (ja3resp.Ja3Hash != options.Ja3Hash) {
			t.Fatal("Expected {} Got {} for Ja3Hash",  options.Ja3Hash, ja3resp.Ja3Hash )
		}
		if (ja3resp.Ja3 != options.Ja3) {
			t.Fatal("Expected {} Got {} for Ja3",  options.Ja3, ja3resp.Ja3 )
		}
		if (ja3resp.UserAgent != options.UserAgent) {
			t.Fatal("Expected {} Got {} for UserAgent",  options.UserAgent, ja3resp.UserAgent )
		}
		
		

		response, err = client.Do("https://http2.pro/api/v1", cycletls.Options{
			Ja3: options.Ja3,
			UserAgent: options.UserAgent,
				
		}, "GET");
		if (response.Response.Status != options.HTTPResponse) {
			t.Fatal("Expected Result Not given")
		} else {
			log.Println("http2: ", response.Response.Status)
		}
	}
}