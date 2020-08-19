#!/bin/bash

openssl req  -new  -newkey rsa:2048  -nodes  -keyout localhost.key  -out localhost.csr
openssl  x509  -req  -days 365  -in localhost.csr  -signkey localhost.key  -out localhost.crt
