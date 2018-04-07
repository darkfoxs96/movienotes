<!doctype html>
<html lang="ru">
  <head>
    {{template "detail/header.tpl" .}}
  </head> 
    
  <body>
     <div id="body_size">
          {{template "detail/navbar.tpl" .}}
          <div id="body_content">
              {{template "body" .}}
          </div>
          <h3 style="display: none" id="converter"></h3>
     </div>
      {{template "detail/footer.tpl" .}}
      {{template "detail/footerDoc.tpl" .}}
  </body>

</html>