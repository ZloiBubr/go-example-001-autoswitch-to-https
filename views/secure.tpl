<!DOCTYPE html>

<html>
<head>
  <title>Main App Title</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="shortcut icon" href="../static/img/document.png" type="image/x-icon" />
  <link rel="stylesheet" href="../static/css/main.css"/>
</head>

<body>
  <header>
    <h1 class="logo">Secure HTTPS page</h1>
    <div class="description">
      This is very secure information.
    </div>
  </header>
  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
    <div>
      Go to Login page:
      <a href="http://{{.WebsiteHTTPS}}">Login Page Link</a>
    </div>
  </footer>
  <div class="backdrop"></div>
</body>
</html>
