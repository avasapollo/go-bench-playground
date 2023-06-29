export ver=v2 && \
  go test -run '^$' -bench '^BenchmarkCompare_with2M$' -benchtime 10s -count 5 \
    -cpu 4 \
    -benchmem \
    -memprofile=bench-result/${ver}.mem.prof -cpuprofile=bench-result/${ver}.cpu.pprof \
    | tee bench-result/${ver}.txt