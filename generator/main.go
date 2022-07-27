package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
)

const latestUrl = "https://releases.mondoo.com/mondoo/latest.json?ignoreCache=1"

var versionMatcher = regexp.MustCompile(`mondoo\/(\d+.\d+.\d+)\/mondoo`)

// Usage: go run main.go
// Example: go run generator/main.go /path
func main() {
	path := "."
	if len(os.Args) == 2 {
		path = os.Args[1]
	}
	log.Printf("Generating files into %s", path)

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

	// write PKGBUILD
	buf := new(bytes.Buffer)
	renderPkgBuild(pb, buf)
	err = os.WriteFile(path+"/PKGBUILD", buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	// write .SRCINFO
	buf2 := new(bytes.Buffer)
	renderScrInfo(pb, buf2)
	err = os.WriteFile(path+"/.SRCINFO", buf2.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

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

var pkgBuildTemplate = `# Maintainer: Mondoo Inc <hello@mondoo.com>
# Maintainer: Dominik Richter <dom@mondoo.com>
# Maintainer: Patrick Münch <patrick@mondoo.com>
#
# TODO:
# - replace the html license with a proper TXT version
# - use upstream oss-licenses instead of bundling it

pkgname=mondoo
pkgver={{ .Version }}
pkgrel=1
pkgdesc="Infrastructure search, analytics, and security analysis"
url="https://mondoo.com"
license=('custom')
source=(
    "https://releases.mondoo.com/mondoo/${pkgver}/mondoo_${pkgver}_linux_amd64.tar.gz"
    'LICENSE.html'
    'OSS-LICENSES.tar.xz'
    'mondoo.service'
    'mondoo.sh'
)
arch=('x86_64')

sha256sums=('{{ .Sha256 }}'
            'c8d346670913c91bf712405e57c2311e6fbda37261f8abfadf9ca7e5fdd768bd'
            'cd99e204a986af5a91f46c43478b28f556a4f50fd9721844d0b600d45ac43cb8'
	        '2febf46353886823e6a61ca15c73e651d71d45579b0a1a17e18905a61387e7e6'
            '92ceefe40c2963f96d02e36743338599cfa9a062d00a5e38580370099b01066c')


package() {
  install -dm755 ${pkgdir}/opt/$pkgname/bin \
                 ${pkgdir}/usr/bin

  cp ${srcdir}/$pkgname ${pkgdir}/opt/$pkgname/bin/.

  install -Dm 755 mondoo.sh ${pkgdir}/usr/bin/mondoo
  install -Dm 644 LICENSE.html "$pkgdir"/usr/share/licenses/$pkgname/LICENSE.html
  install -Dm 644 OSS-LICENSES.tar.xz "$pkgdir"/usr/share/licenses/$pkgname/OSS-LICENSES.tar.xz
  install -Dm 644 mondoo.service "$pkgdir"/usr/lib/systemd/system/mondoo.service
}

#vim: syntax=sh`

func renderPkgBuild(b PkgBuild, out io.Writer) error {
	t := template.Must(template.New("formula").Parse(pkgBuildTemplate))
	return t.Execute(out, b)
}

var pkgSourceInfoTemplate = `pkgbase = mondoo
pkgdesc = Infrastructure search, analytics, and security analysis
pkgver = {{ .Version }}
pkgrel = 1
url = https://mondoo.com
arch = x86_64
license = custom
source = https://releases.mondoo.com/mondoo/{{ .Version }}/mondoo_{{ .Version }}_linux_amd64.tar.gz
source = LICENSE.html
source = OSS-LICENSES.tar.xz
source = mondoo.service
source = mondoo.sh
sha256sums = {{ .Sha256 }}
sha256sums = c8d346670913c91bf712405e57c2311e6fbda37261f8abfadf9ca7e5fdd768bd
sha256sums = cd99e204a986af5a91f46c43478b28f556a4f50fd9721844d0b600d45ac43cb8
sha256sums = 2febf46353886823e6a61ca15c73e651d71d45579b0a1a17e18905a61387e7e6
sha256sums = 92ceefe40c2963f96d02e36743338599cfa9a062d00a5e38580370099b01066c

pkgname = mondoo
`

func renderScrInfo(b PkgBuild, out io.Writer) error {
	t := template.Must(template.New("formula").Parse(pkgSourceInfoTemplate))
	return t.Execute(out, b)
}
