#!/bin/bash -eu

for f in $(find test -type f); do
  case $f in
    *.encrypt)
      continue ;;
    *)
      echo "encrypt $f..."
      f2=$(sed "s|test/||" <<< $f)
      curl -s "localhost/get/$f2?mode=encrypt" > $f.encrypt
    ;;
  esac
done