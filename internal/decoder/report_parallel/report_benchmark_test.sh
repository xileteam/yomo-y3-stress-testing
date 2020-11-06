#!/usr/bin/env bash

set -e

temp_file="../../../docs/temp.out"
report_file="../../../docs/report.out"
go test -bench=. -benchtime=3s -benchmem -run=none | grep Benchmark > ${temp_file} \
  && echo 'finished bench' \
  && cat ${temp_file} \
  && cat ${temp_file} | awk '{print $1,$3}' | awk -F "_" '{print $2,$3"-"substr($4,1,3),substr($4,7)}' | awk -v OFS=, '{print $1,$2,$3}' > ${report_file} \
  && echo 'finished analyse' \
  && cat ${report_file}
