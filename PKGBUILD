# Maintainer: Mondoo Inc <hello@mondoo.com>
# Maintainer: Dominik Richter <dom@mondoo.com>
#
# TODO:
# - replace the html license with a proper TXT version
# - use upstream oss-licenses instead of bundling it

pkgname=mondoo
pkgver=5.36.1
pkgrel=1
pkgdesc="Infrastructure search, analytics, and security analysis"
url="https://mondoo.com"
license=('custom')
source=(
    "https://releases.mondoo.com/mondoo/${pkgver}/mondoo_${pkgver}_linux_amd64.tar.gz"
    'LICENSE.html::https://mondoo.com/terms'
    'OSS-LICENSES.txt'
    'mondoo.service'
    'mondoo.sh'
)
arch=('x86_64')

sha256sums=('e5c1de90449ad58c6621d1b7d929de5279c83b8c05e945f9a693e82caab45dc2'
            '77e73e231a13c3c072e4c8a4812779d34b9bdf9cc6a495f82fd0efc92f07c1aa'
            '8815d47139e8ea19ac407593b91ca232fbfba92bce694b72db159fa53367bb82'
            '2febf46353886823e6a61ca15c73e651d71d45579b0a1a17e18905a61387e7e6'
            '92ceefe40c2963f96d02e36743338599cfa9a062d00a5e38580370099b01066c')


package() {
  install -dm755 ${pkgdir}/opt/$pkgname/bin \
                 ${pkgdir}/usr/bin

  cp ${srcdir}/$pkgname ${pkgdir}/opt/$pkgname/bin/.

  install -Dm 755 mondoo.sh ${pkgdir}/usr/bin/mondoo
  install -Dm 644 LICENSE.html "$pkgdir"/usr/share/licenses/$pkgname/LICENSE.html
  install -Dm 644 OSS-LICENSES.txt "$pkgdir"/usr/share/licenses/$pkgname/OSS-LICENSES.txt
  install -Dm 644 mondoo.service "$pkgdir"/usr/lib/systemd/system/mondoo.service
}

#vim: syntax=sh
