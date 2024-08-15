# GoWeb
Rebuild of my website irlalan.net from an entierly .net C# based web-app to a golang web-app.

Written using vanilla golang for the backend and golang + [templ](https://github.com/a-h/templ) for the frontend with bootstrap Ui.

## Reason for rewrite

After making the website in C# as a way to learn the language better and get to grips (somewhat) with the .net landscape specifically Blazor for webapp 
development I realised I didnt like how clunky and dependancy heavy the process of creating a relativly small scale website using these technologies was.

So after mulling over my feelings I decided to use this as an opportunity to rewrite my website and learn about golang whilst doing so as I was already intrested in learning it and heard 
that it was a simple language to learn and use with great tooling for these kinds of projects.

## Dependancies used

golang, [net/http](https://pkg.go.dev/net/http),  [templ](https://github.com/a-h/templ)

# TOOD:

- Refactor internal url names
- Add Tests
- Refactor Page builders i.e. chaning funcs like PageHeadTemplate to more robust functions
- (Potential) Add markdown to html support so faster blog creations
- Auto gen and update server when blog posted/page edited

