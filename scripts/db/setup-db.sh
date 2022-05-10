#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT"

./mysql.sh
./postgres.sh
