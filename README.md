# Arch Linux AUR PKGBUILD for Mondoo

This repository holds the PKGBUILD go generator and the PKGBUILD files from [https://aur.archlinux.org/packages/mondoo/](https://aur.archlinux.org/packages/mondoo/)

## Publish via aurpublish (reminder)

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

go run ./generator/main.go > mondoo/PKGBUILD

git ap

git cm "<version number"

aurpublish mondoo
```

Powered by [aurpublish](https://github.com/eli-schwartz/aurpublish)