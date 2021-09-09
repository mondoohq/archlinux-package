package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
)

const latestUrl = "https://releases.mondoo.io/mondoo/latest.json?ignoreCache=1"

var versionMatcher = regexp.MustCompile(`mondoo\/(\d+.\d+.\d+)\/mondoo`)

// Usage: go run main.go
// Example: go run generator/main.go > PKGBUILD
func main() {
	latest, err := fetchLatest()
	if err != nil {
		log.Fatal(err)
	}

	// filter by linux and amd64
	pb := PkgBuild{}

	for i := range latest.Files {
		f := latest.Files[i]

		m := versionMatcher.FindStringSubmatch(f.Filename)
		if len(m) == 2 {
			pb.Version = m[1]
		}

		if f.Platform == "linux" && strings.HasSuffix(f.Filename, "amd64.tar.gz") {
			pb.Sha256 = f.Hash
		}
	}

	buf := new(bytes.Buffer)
	renderPkgBuild(pb, buf)
	fmt.Println(buf.String())
	os.Exit(0)
}

type Latest struct {
	Files []File `json:"files"`
}

type File struct {
	Filename string `json:"filename"`
	Size     int    `json:"size"`
	Platform string `json:"platform"`
	Hash     string `json:"hash"`
}

func fetchLatest() (*Latest, error) {
	resp, err := http.Get(latestUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var latest Latest
	if err := json.Unmarshal(data, &latest); err != nil {
		return nil, err
	}

	return &latest, nil
}

type PkgBuild struct {
	Version string `json:"version"`
	Sha256  string `json:"sha256"`
}

var pkgBuildTemplate = `# Maintainer: Mondoo Inc <hello@mondoo.io>
# Maintainer: Dominik Richter <dom@mondoo.io>
#
# TODO:
# - replace the html license with a proper TXT version
# - use upstream oss-licenses instead of bundling it

pkgname=mondoo
pkgver={{ .Version }}
pkgrel=1
pkgdesc="Infrastructure search, analytics, and security analysis"
url="https://mondoo.io"
license=('custom')
source=(
    "https://releases.mondoo.io/mondoo/${pkgver}/mondoo_${pkgver}_linux_amd64.tar.gz"
    'LICENSE.html::https://mondoo.io/terms'
    'OSS-LICENSES.txt'
    'mondoo.service'
    'mondoo.sh'
)
arch=('x86_64')

sha256sums=('{{ .Sha256 }}'
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

#vim: syntax=sh`

func renderPkgBuild(b PkgBuild, out io.Writer) error {
	t := template.Must(template.New("formula").Parse(pkgBuildTemplate))
	return t.Execute(out, b)
}
