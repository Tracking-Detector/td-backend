load("@rules_python//python:defs.bzl", "py_test")

py_test(
    name = "model_test",
    srcs = [
        "model_test.py",
    ],
    visibility = ["//visibility:public"],
    deps = ["//python/td_common/model"],
)

py_test(
    name = "dataset_test",
    srcs = [
        "dataset_test.py",
    ],
    visibility = ["//visibility:public"],
    deps = ["//python/td_common/model"],
)

py_test(
    name = "training_run_test",
    srcs = [
        "training_run_test.py",
    ],
    visibility = ["//visibility:public"],
    deps = ["//python/td_common/model"],
)
