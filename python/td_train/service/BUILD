load("@rules_python//python:defs.bzl", "py_library")

py_library(
    name = "service",
    srcs = [
        "training_data_service.py",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//python/td_common/config",
        "//python/td_common/model",
        "//python/td_common/minio",
        "@pypi//pymongo",
        "@pypi//numpy",
        "@pypi//pandas",
    ],
)