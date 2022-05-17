_default:
    @just -u -l

buf-gen:
    which buf || brew install bufbuild/buf/buf
    buf mod update ./proto
    buf generate