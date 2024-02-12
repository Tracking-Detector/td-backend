load("@rules_python//python:defs.bzl", "py_test")

py_test(
    name = "model_test",
    srcs = [
        "dataset_test.py",
        "model_test.py",
        "training_run_test.py",
    ],
    visibility = ["//visibility:public"],
    deps = ["//python/td_common/model"],
)
