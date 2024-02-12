load("@rules_python//python:defs.bzl", "py_test")
py_test(
    name = "model_test",
    srcs = [
        "model_test.py",
        "training_run_test.py",
        "dataset_test.py",
    ],
    deps = ["//python/td_common/model:model"],
    visibility = ["//visibility:public"],
)