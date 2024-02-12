load("@rules_python//python:defs.bzl", "py_library")

py_library(
    name = "repository",
    srcs = [
        "base_repository.py",
        "dataset_repository.py",
        "model_repository.py",
        "training_run_repository.py",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//python/td_common/config",
        "//python/td_common/model",
        "@pypi//pymongo",
    ],
)
