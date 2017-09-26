<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8"/>
    <title>File uploader</title>
  </head>
  <body>
    <h1>Upload a file</h1>
    <form action="receiver" method="POST" enctype="multipart/form-data">
      <input name="ufile" type="file" />
      <input type="submit" />
    </form>
  </body>
</html>
