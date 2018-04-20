: [Acasa](../index.html) / [Retele de calculatoare](./retele_calculatoare.html)

# Retele de calculatoare

Modelul OSI este o reprezentarea conceptuala a organizarii retelelor de calculatoare.

![Model OSI](../img/osi.png)

## Linux și rețelistica

Linux inițiază tabela de routare a sistemului în timpul pornirii sistemului de operare. După ce kernelul termină secvența de inițializare, acesta rulează anumite utilitare de sistem si verifică fișiere de configurare ale sistemului de operare, prin care stabiliește identitatea calculatorului la nivel de rețea.

Aici vorbim de informații precum propria identitate (adresa IP), inițializarea interfețelor de rețea disponibile pe sistem, setarea rutelor de conectare spre rețele externe (spre internet).

Tot acest proces poate fi static (setarea statică a configurării) sau dinamică (folosirea protocolului DHCP). Pe scurt, am descris unul din rolurile serviciul de rețelistică din Linux.

Vom discuta în continuare de implementarea acestui modul in versiunea Centos 7/Red Hat Enterprise 7, dar aplicabile si pentru distribuții precum Fedora 19.


## Utilitare

Serviciul specific care se ocupă de resursele de rețelistică se numește NetworkManager. Acesta este un serviciu de tip daemon. 

Utilitarele prin care putem interacționa cu acest serviciu sunt :

- nmtui - este un program bazat pe o interfață text disponibil in consolă/terminal

- nmcli - este un utilitar bazat pe comenzi si parameteri, disponibil in consolă/terminal

- control-center - utilitar grafic

- nm-connection-editor - utilitar grafic cu funcții suplimentare


NetworkManager ar trebui sa vina preinstalat, fiind un serviciu esențial, dar se poate instala cu ajutorul comenzii
```bash
user@linux ~ $ yum install NetworkManager
```

Pentru a verifica starea serviciului, putem folosi următoarele comenzi în terminal
```bash
user@linux ~ $ systemctl status NetworkManager
NetworkManager.service - Network Manager
   Loaded: loaded (/lib/systemd/system/NetworkManager.service; enabled)
   Active: active (running) since Fri, 08 Mar 2013 12:50:04 +0100; 3 days ago

### Putem porni serviciul
user@linux ~ $ systemctl start NetworkManager
### Putem opri serviciul
user@linux ~ $ systemctl stop NetworkManager
### Putem reporni serviciul
user@linux ~ $ systemctl restart NetworkManager  
### Putem pune serviciul să pornească automat la inițializarea sistemului
user@linux ~ $ systemctl enable NetworkManager
### Putem pune serviciul să NU pornească automat la inițializarea sistemului
user@linux ~ $ systemctl disable NetworkManager 
```

Informațiile setate cu ajutorul acestor utilitare se pot găsi in locația */etc/sysconfig/* .

Fișierele cu setări specifice conexiunilor de VPN, PPPoE se pot găsi în locația */etc/NetworkManager/* .

Dacă aceste fișiere sunt editate, pentru a se aplica schimbările specificate în fișiere, trebuie notificat serviciul de NetworkManager pentru a reciti fișierele de configurare.   

```bash
### Pentru a reciti toate fișierele 
user@linux ~ $ nmcli connection reload
### Pentru a reciti un fișier specific
user@linux ~ $ nmcli con load /etc/sysconfig/network-scripts/ifcfg-ifname
```

## Firewall

Firewallul este o componentă filtrare a traficului de reţea. El este necesar pentru a ne asigura că vom accepta doar anumite tipuri de conexiuni, pe anumire porturi/protocoale care au fost în prealabil revizuite si aprobate.

În Centos7, acestă componentă se numeşte firewalld - dynamic firewall daemon.

Vom interacţiona cu acest daemon prin intermediul utilitarului *firewall-cmd*

Serviciul firewalld vine cu anumite zone predefinite de firewall.

Putem asocia zonele cu intervale de IP-uri (subreţele) pe care putem aplica anumite filtre/reguli
```bash
root@linux ~ $ firewall-cmd --get-zones
block dmz drop external home internal public trusted work
```

Vom vedea care sunt dispozitivele de reţea (interfeţele) disponibile - enp3s0 (cablu) şi wlp2s0 (wireless) - şi în care zona sunt asociate intervalele de IP.
```bash
### Vom lista interfetele şi reţelele asociate
root@linux ~ $ firewall-cmd --get-active-zone
public
  interfaces: enp3s0 wlp2s0
trusted
  sources: 192.168.1.1/24 192.168.0.1/24
```

Pe o anumită zonă, putem vedea care sunt serviciile (porturile şi protocoalele predefinite) accesibile din acea reţea (zona home)
```bash
### Vom afişa serviciile accesibile asociate cu zona home
root@linux ~ $ firewall-cmd --zone=home --list-services
ssh mdns samba-client dhcpv6-client
```
Pentru mai multe detalii referitoare la serviciile predefinite, puteţi verifica locaţia */usr/lib/firewalld/services* unde puteţi găsi un XML de configurare pentru fiecare serviciu predefinit şi disponibil in firewalld.


Pe o anumită zonă, putem vedea care sunt serviciile (porturile predefinite) accesibile din acea reţea (zona home).
În exemplul de mai jos putem observa că HTTP şi HTTPS vor fi accesibile atunci când o cerere către aceste servicii va veni din zona publică. 
```bash
root@linux ~ $ firewall-cmd --zone=public --list-services
dhcpv6-client http https
```

Vom afişa serviciile accesibile asociate cu zona external - vedem doar SSH ca şi serviciu disponibil.
```bash
root@linux ~ $ firewall-cmd --zone=external --list-all
external
  target: default
  icmp-block-inversion: no
  interfaces:
  sources:
  services: ssh
  ports:
  protocols:
  masquerade: yes
  forward-ports:
  source-ports:
  icmp-blocks:
  rich rules:
```

Pentru a adăuga definiţii noi de servicii - de exemplu myapp - o aplicaţie pe care voi o dezvoltaţi şi care trebuie sa primească conexiuni pe TCP port 1234 şi UDP pe 1245.


```bash
### Vom defini un nou serviciu 
root@linux ~ $ firewall-cmd --permanent --new-service=myapp
success
### In acest serviciu vom defini portul 1234 TCP si 1245 UDP
root@linux ~ $ firewall-cmd --permanent --service=myapp --add-port=1234/tcp
success
root@linux ~ $ firewall-cmd --permanent --service=myapp --add-port=1245/udp
success

### Putem valida ca un nou fişier XML a fost creat, unde vom găsi configurările făcute mai sus
root@linux ~ $ cat /etc/firewalld/services/myapp.xml
<?xml version="1.0" encoding="utf-8"?>
<service>
  <port protocol="tcp" port="1234"/>
  <port protocol="udp" port="1245"/>
</service>

### Vom adăuga nou serviciu myapp în zona publică
### Atenţie! Fără parametrul --permanent această schimbare nu va persista după repornirea sistemului de operare
root@linux ~ $ firewall-cmd --zone=public --permanent --add-service=myapp
success

### Este nevoie de un reload pentru ca schimbările noastre să fie aplicate
root@linux ~ $ firewall-cmd --reload
success

### Putem valida ca avem un nou serviciu acesibil în zona publică
root@linux ~ $ firewall-cmd --zone=public --list-all
public (active)
  target: default
  icmp-block-inversion: no
  interfaces: enp3s0 wlp2s0
  sources:
  services: dhcpv6-client http https myapp
  ports: 444/tcp
  protocols:
  masquerade: no
  forward-ports:
  source-ports:
  icmp-blocks:
  rich rules:


```
# Exercitii
1. Verificam statusul serviciul de Firewall si il pornim, daca nu este inca pornit, pe VM-ul Proxy
2. Verificam lista de servicii accesibile, pe VM-ul Proxy
3. Stergem serviciul "dhcpv6-client" din lista de servicii accesibile prin firewall, pe VM-ul Proxy
4. Adaugam serviciile HTTP si HTTPS in lista de servicii accesibile prin firewall, pe VM-ul Proxy.
5. Pe VM-ul Web, vom adauga portul 18080 si protocolul TCP in firewall.

#### [Instalarea de software](./instalare_software.html)

* * *
![license](https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png)

This work is licensed under a [Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License](http://creativecommons.org/licenses/by-nc-sa/4.0/)