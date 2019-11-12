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

## Post Comments

```
<form name="comento">
  <label>Name:<br><input type="text" name="name" value="John Doe"></label><br>
  <br>
  <label>Comment:<br><textarea name="comment"></textarea></label><br>
  <br>
  <button type="button" onclick="post()">Submit</button>
</form>

<script>
function post() {
  let action;
  if (/\/$/.test(location.pathname)) {  // Does the URL end with "/"?
    action = 'https://your.domain/post?thread=' + location.pathname + 'index.html';
  } else {
    action = 'https://your.domain/post?thread=' + location.pathname;
  }
  document.comento.method = 'post';
  document.comento.action = action;
  document.comento.submit();
}
</script>
```

## Get Comments

Data Form is JSON.

```
# Get comments for a page.
$ curl https://your.domain/get?thread=/page1.html
[{"name":"name1","comment":"comment1"},{"name":"name2","comment":"comment2"}]

# Do you use staging level certificates?
$ curl -k https://your.domain/get?thread=/page1.html


# Get all comments.
$ curl https://your.domain/get_all
[{"thread":"/page1.html","name":"name1","comment":"comment1"},{"thread":"/page1.html","name":"name2","comment":"comment2"},{"thread":"/page2.html","name":"name1","comment":"comment1"}]

# Do you use staging level certificates?
$ curl -k https://your.domain/get_all
```

Comento's comment data is in "comento/comento.db".  
You can open it with "$ sqlite3 comento/comento.db".

Comento uses cccc as a reverse proxy.  
See "https://github.com/taku-n/cccc".
