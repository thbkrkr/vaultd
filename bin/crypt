#!/bin/bash -eu

encrypt() {
  for f in $(find test -type f); do
    case $f in
      *.encrypt)
        continue ;;
      *)
        echo "encrypt $f..."
        f2=$(sed "s|test/||" <<< $f)
        curl -s "localhost:4242/get/$f2?mode=encrypt" > $f.encrypt
      ;;
    esac
  done
}

decrypt() {
  for f in $(find test -type f); do
    case $f in
      *.encrypt)
        f2=$(sed "s|test/||" <<< $f)
        f2=$(sed "s|.encrypt||" <<< $f2)
        f=$(sed "s|.encrypt||" <<< $f)
        echo "decrypt $f..."
        curl -s "localhost:4242/get/$f2" > $f
      ;;
      *)
        continue ;;
    esac
  done
}

main() {
  case "${1:-}" in
    d|decrypt) decrypt ;;
    *)  encrypt ;;
  esac
}

main "$@"