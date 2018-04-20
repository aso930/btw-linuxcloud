: [Acasa](../index.html) / [Configurarea serviciilor](./configurarea_serviciilor.html)

# Configurarea serviciilor

In Linux, configurarea aplicatiilor, serviciilor si a sistemului se face exclusiv prin editarea de fisiere text de configurarea. Conform practicilor, aceste fisiere se regasesc in directorul `/etc`. Aceste fisiere pot fi editate manual, sau pot fi schimbate folosind unelte oferite de dezvoltatorii aplicatiei. Cum sunt editate difera de la aplicatie la aplicatie. 

Serverul web Apache are fisierele de configurare in directorul `/etc/httpd/conf`. Fisierul principal de configurare se numeste `httpd.conf`. In acest fisier putem configura toate aspectele ce tin de functionarea server-ului web. Pe noi ne intereaseaza sa configuram server-ul web sa functioneaza ca un proxy pentru aplicatia noastra.

Fiecare aplicatie are formatul ei pentru fisierul de configurare. Este foarte important ca atunci cand vreti sa configurati o aplicatie in Linux, sa cititi foarte atent documentatia, deoarece nu exista un format universal.

Pe langa fisierul principal de configurare pentru serverul Apache, `httpd.conf`. Apache cauta configurari suplimentare in directorul `/etc/httpd/conf.d`. Este indicat pentru o buna organizare si separare a fisierelor de configurare sa cream un fisier nou in conf.d pentru configurarea de proxy. Chiar daca noi avem un config foarte simplu, daca aveam foarte multe website-uri definite, e mult mai usor sa ai un fisier separat de configurare pentru fiecare, decat sa ai toata configurarea in `httpd.conf`.

Vom crea fisierul `btw-web.conf` in `/etc/httpd/conf.d` cu continutul urmator:

```xml
<VirtualHost *:80>
ProxyPreserveHost On
ProxyPass / http://192.168.111.2:18080/
ProxyPassReverse / http://192.168.111.2:18080/
</VirtualHost>
```

Pentru a putea folosi functia de proxy a httpd, trebuie sa ii dam permisiunea respectiva:
```bash
setsebool -P httpd_can_network_connect 1
```

# Exercitii
1. Vom configura serviciul httpd pe VM-ul Proxy, pentru a asculta requesturi HTTP.
2. Verificam pagina de test httpd.
3. Vom configura o locatie separata care sa ne trimita spre VM-ul Web.

#### [Pornirea aplicatiilor](./pornirea_aplicatiilor.html)

* * *
![license](https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png)

This work is licensed under a [Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License](http://creativecommons.org/licenses/by-nc-sa/4.0/)