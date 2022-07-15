# Arch Linux AUR PKGBUILD for Mondoo

This repository holds the PKGBUILD go generator and the PKGBUILD files from [https://aur.archlinux.org/packages/mondoo/](https://aur.archlinux.org/packages/mondoo/)

## Setup aurpublish

``` bash
mkdir .git/hooks/
aurpublish setup
```

- add to /home/user/.ssh/config

``` text
Host aur aur.archlinux.org
        User aur
        Hostname aur.archlinux.org
        IdentityFile ~/.ssh/arch-linux-aur
```

``` bash
aurpublish -p mondoo
```

## Update aur package

``` bash
make
```

``` bash
git add mondoo/PKGBUILD

git commit -s -m "<version number"

aurpublish mondoo
```
