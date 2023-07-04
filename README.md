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