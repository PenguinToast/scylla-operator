#!/usr/bin/env bash

set -euExo pipefail
shopt -s inherit_errexit

url="$1"

tmp="$( mktemp )"
trap 'rm -f ${tmp}' EXIT

curl --fail --retry 5 -L "${url}" > "${tmp}"

cat "${tmp}"
