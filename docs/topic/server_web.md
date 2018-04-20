: [Acasa](../index.html) / [Server web](./server_web.html)

# Server web

Un server web este o aplicatie care primeste si proceseaza request-uri facute folosind protocolul HTTP.

HTTP care faciliteaza existenta web-ului asa cum il cunoastem azi. Prin HTTP aplicatiile pot comunica si transfera informatii.

La nivelul cel mai simplu comunicarea in HTTP se face prin apelarea unor metode predefinite si coduri de status. Un client apeleaza metode pe server, cu datele pe care vrea sa le transmita atasate in `request body`. La primirea request-ului, serverul web il va procesa si va trimite inapoi un status code si datele cerute in `response body` 

Metodele HTTP sunt o serie de verbe predifinite, pe baza carora serverul web va executa anumite operatiuni. Cele mai intalnite metode HTTP sunt GET, POST si PUT.

- GET este folosit pentru a cere o resursa de pe server, spre exemplu pentru a vedea o pagina web
- POST este folosit pentru a trimite date catre serverul web spre prelucrare. Datele sunt stocate in `request body`
- PUT este folosit pentru a modifica resursele identificate prin path-ul request-ului.

Cum arata un request si response simplu in http:

```http
* Connected to localhost (::1) port 18080 (#0)
> GET / HTTP/1.1
> User-Agent: curl/7.29.0
> Host: localhost:18080
> Accept: */*
>

< HTTP/1.1 200 OK
< Date: Fri, 20 Apr 2018 09:08:12 GMT
< Content-Length: 692
< Content-Type: text/html; charset=utf-8
<
<!DOCTYPE html>
<html>
<head>
<title>Linux and the cloud - Web App</title>
(...)

```

Codurile de status sunt mijloacele prin care protocolul HTTP categorizeaza un raspuns. Acestea sunt standardizate si contin un cod de 3 cifre (e.g. 404) si un mesaj (e.g. Not Found). Pe baza lor ne putem da seama daca request-ul a mers sau nu si avem si un indiciu de ce.

Exista 5 categorii mari de coduri de status:

 - 1xx - mesaje informative
 - 2xx - mesaje de succes
 - 3xx - mesaje de redirectionare
 - 4xx - mesaje de eroare pe client
 - 5xx - mesaje de eroare pe server

# Exercitii
1. Vom deschide un browser si vom activa untilitarul de diagnoza prin tasta F12.
2. In utilitarul de diagnoza vom activa tab-ul Network, pentru a monitoriza requesturile HTTP.
3. Vom investiga un request HTTP catre un site - [TEST](http://status.aso.re/)


#### [Configurarea serviciilor](./configurarea_serviciilor.html)

* * *
![license](https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png)

This work is licensed under a [Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License](http://creativecommons.org/licenses/by-nc-sa/4.0/)