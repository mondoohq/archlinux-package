# Arch Linux AUR PKGBUILD for Mondoo

This repository holds the PKGBUILD go generator and the PKGBUILD files from [https://aur.archlinux.org/packages/mondoo/](https://aur.archlinux.org/packages/mondoo/)


## Install cnquery and cnspec

The packages are published on [AUR](https://aur.archlinux.org):

- [cnquery](https://aur.archlinux.org/packages/cnquery)
- [cnspec](https://aur.archlinux.org/packages/cnspec)

Use one of the [AUR_helpers](https://wiki.archlinux.org/title/AUR_helpers) like `yay` to install the packages:

```
# install cnquery
yay -Ss cnquery

# install cnspec
yay -Ss cnspec
```

## Test github action

- create the `.secret` file with the following content:

```
AUR_USERNAME="Patrick Münch"
AUR_EMAIL="patrick@mondoo.com"
AUR_SSH_PRIVATE_KEY="-----BEGIN OPENSSH PRIVATE KEY-----\n....\n-----END OPENSSH PRIVATE KEY-----\n"
```

- create the `sample-event.json` file with the following content:

```
{
  "action": "workflow_dispatch",
  "inputs": {
      "version": "6.13.1"
  }
}
```

- run the following command:

```bash
act -j aur-publish --secret-file .secrets --eventpath sample-event.json -v
```