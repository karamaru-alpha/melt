#!/bin/bash

skeema diff local -p$MYSQL_ROOT_PASSWORD --allow-unsafe --skip-verify

read -p "破壊的にmigrateしますか？ (y/n) :" YN
if [ "${YN}" = "y" ]; then
  skeema push local -p$MYSQL_ROOT_PASSWORD --allow-unsafe --skip-verify
fi
