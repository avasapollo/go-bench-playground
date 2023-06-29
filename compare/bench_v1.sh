export ver=v1 && \
  go test -run '^$' -bench '^BenchmarkSum_with2M$' -benchtime 10s -count 6 \
    -cpu 4 \
    -benchmem \
    -memprofile=bench-result/${ver}.mem.prof -cpuprofile=bench-result/${ver}.cpu.pprof \
    | tee bench-result/${ver}.txt