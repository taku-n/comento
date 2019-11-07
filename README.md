# cccc
Continuous enCryption, Continuous Certification. Using external volume to keep certification.

Example
=======

You get a reverse proxy with TLS and a web server.

```
$ git clone https://github.com/taku-n/cccc.git  
$ cd cccc  
$ nvim docker-compose.yml  
# Set environment variables.  
# Choose STAGE=staging for --test-cert or STAGE=production for not.

$ docker-compose up -d
```
