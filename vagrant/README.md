# Cum pornim masinile virtuale

1. Descarcam si instalam Virtual Box de [aici](https://download.virtualbox.org/virtualbox/5.2.8/VirtualBox-5.2.8-121009-Win.exe)

2. Descarcam si instalam Vagrant de [aici](https://releases.hashicorp.com/vagrant/2.0.3/vagrant_2.0.3_x86_64.msi)

3. Copiem fisierul denumit Vagrantfile (pe care il gasiti aici) pe calculatorul personal, intr-un folder separat.

4. Deschidem Command Prompt in folderul unde am descarcat fisierul Vagrantfile. Apasam pe bara de adrese in File explorer, scriem `cmd` si apasam tasta `Enter`.

5. Ne asiguram ca suntem in folderul unde e fisierul Vagrantfile. Scriem `dir` in fereastra Command Prompt, ar trebui sa vedem fisierul Vagrantfile. Daca nu apare, verificati ca sunteti in folderul care trebuie.

6. Executati comanda `vagrant up`

7. Lasati comanda sa ruleze. Aceasta va dura cateva minute, in functie de viteza conexiunii la internet. La final va aparea ca s-a executat cu succes.

8. Opriti masinile virtuale folosind `vagrant halt` si inchideti fereastra