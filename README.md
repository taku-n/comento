# Comento
A commenting server for static pages.

Example
=======

You need a server with Docker and its domain name.

```
$ git clone https://github.com/taku-n/comento.git
$ cd comento
$ nvim docker-compose.yml
# Set environment variables.
# Choose STAGE=staging for --test-cert or STAGE=production for not.

$ docker-compose up -d
```

Comento's comment data is in "comento/comento.db".
You can open it with "$ sqlite3 comento/comento.db".
