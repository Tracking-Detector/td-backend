load("@rules_python//python:defs.bzl", "py_test")

py_test(
    name = "training_data_service_test",
    srcs = [
        "training_data_service_test.py",
    ],
    data = [":testdata/expected_encoding.csv.gz"],
    visibility = ["//visibility:public"],
    deps = [
        "//python/td_train/service",
        "//python/td_common/model",
        "//python/td_common/minio"
    ],
)

