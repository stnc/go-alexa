# my-go-alexa
//bu kod puplic key i cikarir
//openssl x509 -pubkey -noout -in private.pem >> public.key


sudo openssl dgst -sha1 -sign private.pem -out /tmp/selman.sha1 selman.sha1
sudo openssl base64 -in /tmp/selman.sha1 -out signature.sha1
rm /tmp/selman.sha1


#verify

sudo openssl base64 -d -in signature.sha1 -out /tmp/selman.sha1
sudo openssl dgst -sha1 -verify public.key -signature /tmp/selman.sha1 selman.sha1
rm /tmp/selman.sha1

Indent = girinti prononce inndent 
intent = niyet - amac gayet intent



utterance = tr = soyleyis ifade bicimi , ses cikarma  adIrIns speak
slot = yuva , yarik, yerlestirmek sLat

go test . -v   
go test . -coverprofile=cover.out    
go tool cover -func=cover.out    
go tool cover -html=cover.out


https://www.golangprograms.com/go-language/struct.html  https://kevin-yang.medium.com/golang-embedded-structs-b9d20aadea84
https://www.golangprograms.com/go-language/interface.html -- lesson work ? -> Interface Accepting Address of the Variable == this example

https://github.com/waringer/Alexa-Radio --bundan model skil falan olsuturma ornekleri al

https://github.com/go-alexa/alexa/blob/master/events/events.go#L64

https://github.com/jsgoecke/lambda-go  js json kodlar var

https://github.com/ericdaugherty/alexa-skills-kit-golang/blob/master/alexa_test.go   test yazma ornegi
https://github.com/nraboy/alexa-slick-dealer/blob/master/main_test.go bu da test

https://github.com/go-alexa/alexa/tree/master  burada da test var

https://github.com/go-alexa/alexa/blob/master/server/server_test.go   buradaki function icinde funtion teknigini mutlaka ogren

go ile rouer yazmak https://benhoyt.com/writings/go-routing/ ve https://golangdocs.com/golang-mux-router  konu hakkinda iyisi


//https://www.tutorialspoint.com/golang-program-that-uses-structs-as-map-keys


//https://github.com/NerdCademyDev/golang  burada generic var
//https://gosamples.dev/enum/ https://articles.wesionary.team/working-with-constants-and-iota-in-golang-460c64792d40  const ile auto inc

https://tutorialedge.net/golang/improving-your-tests-with-testify-go/
# test yazmak
https://github.com/patst/alexa-skills-kit-for-go/tree/master/alexa  bundan basla hazir seyler var
https://articles.wesionary.team/writing-unit-tests-with-testify-84b859dcf91


https://speedscale.com/blog/testing-golang-with-httptest/  
github.com/smartystreets/gunit

https://github.com/alexa/alexa-skills-kit-sdk-for-nodejs/blob/2.0.x/ask-sdk-express-adapter/tst/verifier/index.spec.ts
https://github.com/alexa/alexa-skills-kit-sdk-for-nodejs/blob/2.0.x/ask-sdk-express-adapter/lib/verifier/index.ts



https://github.com/jakubsuchy/amazon-alexa-php
https://github.com/arienmalec/alexa-go/blob/master/request_test.go
https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742#.mjytgulbg
//https://github.com/nraboy/alexa-slick-dealer/blob/master/main.go
https://gitlab.com/go-box/pongo2gin/-/blob/master/v4/render_test.go
https://github.com/nraboy/alexa-slick-dealer/blob/master/main_test.go
https://github.com/brentnd/go-snowboy/blob/master/snowboy_test.go
https://github.com/arienmalec/alexa-go/blob/master/request_test.go
https://github.com/b00gizm/golexa/blob/master/event_test.go
https://github.com/jackmcguire1/alexa-chatgpt/blob/main/internal/api/handler_test.go
https://www.youtube.com/watch?v=YkghFAxdoyU&t=833s

go kucuk kodlamar ornegi
https://github.com/ericdaugherty/alexa-skills-kit-golang
https://github.com/go-alexa/alexa
https://github.com/patst/alexa-skills-kit-for-go/tree/master

https://github.com/mikeflynn/go-alexa/tree/master  most stars

for devices
https://github.com/webability-go/alexa

// import 	"github.com/aivahealth/goalexa/alexaapi"
//https://github.com/nraboy/alexa-slick-dealer/blob/master/main.go# goalexa


TODO
https://github.com/patst/alexa-skills-kit-for-go/blob/master/alexa/http_test.go  buradaki valid request json falan eklenecek yayindan sonra 




https://developer.amazon.com/en-US/docs/alexa/custom-skills/speech-synthesis-markup-language-ssml-reference.html
https://github.com/mikeflynn/go-alexa/blob/master/skillserver/ssml-builder.go
https://github.com/webability-go/alexa bunun yapisi da iyi gibi
https://github.com/go-alexa/alexa/blob/master/server/server_test.go  sunun yaptigi gibi chain olayini arka arkaya donusturmek
https://github.com/ericdaugherty/alexa-skills-kit-golang/tree/master/samples/helloworld aws de nasil olur
https://github.com/DasJott/alexa-sdk-go
https://github.com/skillkit/go-alexa
https://github.com/stnc/chi v1
https://github.com/aivahealth/goalexa/blob/master/attributes.go  bunun icin generic ogrenmek lazim yapilacak
https://webflow.com/pricing  hazir site yapan arac 
https://github.com/bxcodec/go-clean-arch/tree/master  benim projemi bunun gibi parcalara ayirabilirim VERRRYYYYYYYY IMPORTANT
https://github.com/patst/alexa-skills-kit-for-go/blob/master/alexa/http_test.go#L87  bunun uzerinde calis cok onemli cok cok VERRRYYYYYYYY IMPORTANT


ssl
https://www.tweaking4all.com/network-internet/create-self-signed-ssl-certificate/
https://www.baeldung.com/linux/openssl-extract-certificate-info