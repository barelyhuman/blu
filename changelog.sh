# #!/bin/bash
version=v0.0.4-dev.1

# Add check for wget
if ! command -v wget &> /dev/null
then
    brew install wget
fi

curl -s https://api.github.com/repos/barelyhuman/commitlog/releases/tags/$version \
| grep "browser_download_url.*commitlog[^extended].*darwin-amd64\.tar\.gz" \
| cut -d ":" -f 2,3 \
| tr -d \" \
| wget -qi -

tarball="$(find . -name "*darwin-amd64.tar.gz*")"
if test -z "$tarball" 
then
      echo "failed to get commitlog, skipping...." || true
      echo "" > CHANGELOG.txt
else
    tar -xzf $tarball
    chmod +x commitlog
    ./commitlog -i fix,feat,refactor > CHANGELOG.txt
fi
