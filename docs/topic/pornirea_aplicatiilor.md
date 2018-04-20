: [Acasa](../index.html) / [Pornirea aplicatiilor](./pornirea_aplicatiilor.html)

# Pornirea aplicatiilor

Pentru a putea porni aplicatii in Linux, prima oara trebuie sa ne asiguram ca fisierul are bitul de executie setat. Putem face asta prin comanda `ls -l` si inspectarea permisiunilor. Daca avem `x` setat pentru utilizator atunci acel fisier poate fi executat.

Trebuie sa stiti ca atunci cand executati o aplicatie din terminalul pe care v-ati conectat cu ssh, aceasta aplicatie este legata de procesul de `bash` din care a fost executata, daca inchidem sesiunea de ssh, se va inchide si aplicatia deschisa. In mod evident asta nu este ce vrem sa se intample cand dorim sa pornim un serviciu care sa fie mereu deschis.

Pentru a obtine comportamentul dorit, trebuie sa configuram aplicatia noastra web ca serviciu, astfel ea pornirea si oprirea ei fiind administrata de sistem.

Exista mai multe metode prin care un serviciu poate fi definit, in functie de distributia de Linux folosita. Familia de distributii din care face CentOS foloseste `systemd` pentru a defini servicii. Un serviciu se defineste prin crearea unui fisier de servicii in `/usr/lib/systemd/system`. Noi vom crea serviciul `btw-web.service`.

```bash
[Unit]
Description=BTW Web Application
After=network.target

[Service]
Type=simple
PrivateTmp=true
ExecStart=/home/btwlinux/btw
User=btwlinux
Group=btwlinux

[Install]
WantedBy=multi-user.target
```

* * *
![license](https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png)

This work is licensed under a [Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License](http://creativecommons.org/licenses/by-nc-sa/4.0/)