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


test yazmak
https://github.com/jakubsuchy/amazon-alexa-php
https://github.com/arienmalec/alexa-go/blob/master/request_test.go
https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742#.mjytgulbg
//https://github.com/nraboy/alexa-slick-dealer/blob/master/main.go



go kucuk kodlamar ornegi
https://github.com/ericdaugherty/alexa-skills-kit-golang
https://github.com/go-alexa/alexa
https://github.com/patst/alexa-skills-kit-for-go/tree/master

https://github.com/mikeflynn/go-alexa/tree/master  most stars

for devices
https://github.com/webability-go/alexa

