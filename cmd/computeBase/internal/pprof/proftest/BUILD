# Description:
#   Auto-imported from github.com/zdyj3170101136/pprofutil/cmd/computeBase/internal/pprof/proftest

licenses(["notice"])

package(
    default_applicable_licenses = ["//third_party/golang/pprof:license"],
    default_visibility = ["//third_party/golang/pprof/internal:friends"],
)

go_library(
    name = "proftest",
    srcs = ["proftest.go"],
    embedsrcs = ["testdata/large.cpu"],
)
