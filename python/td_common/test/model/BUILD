load("@rules_python//python:defs.bzl", "py_test")

py_test(
    name = "model_test",
    srcs = [
        "model_test.py",
    ],
    tags = ["unittest"],
    visibility = ["//visibility:public"],
    deps = ["//python/td_common/model"],
)

py_test(
    name = "dataset_test",
    srcs = [
        "dataset_test.py",
    ],
    tags = ["unittest"],
    visibility = ["//visibility:public"],
    deps = ["//python/td_common/model"],
)

py_test(
    name = "training_run_test",
    srcs = [
        "training_run_test.py",
    ],
    tags = ["unittest"],
    visibility = ["//visibility:public"],
    deps = ["//python/td_common/model"],
)
