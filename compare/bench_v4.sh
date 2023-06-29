export ver=v4 && \
  go test -run '^$' -bench '^BenchmarkSCompare$' -benchtime 10s -count 8 \
    -cpu 4 \
    -benchmem \
    -memprofile=bench-result/${ver}.mem.prof -cpuprofile=bench-result/${ver}.cpu.pprof \
    | tee bench-result/${ver}.txt