#!/bin/bash

set -euo pipefail
cd "$(dirname "$0")"

goimports -w -local github.com/pusrenk .
