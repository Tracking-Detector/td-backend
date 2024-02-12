load("@rules_python//python:defs.bzl", "py_test")

py_test(
    name = "model_repository_test",
    srcs = [
        "model_repository_test.py",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//python/td_common/model",
        "//python/td_common/repository",
        "@pypi//pymongo",
        "@pypi//pytest",
        "@pypi//testcontainers",
    ],
)

py_test(
    name = "dataset_repository_test",
    srcs = [
        "dataset_repository_test.py",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//python/td_common/model",
        "//python/td_common/repository",
        "@pypi//pymongo",
        "@pypi//pytest",
        "@pypi//testcontainers",
    ],
)

py_test(
    name = "training_run_repository_test",
    srcs = [
        "training_run_repository_test.py",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//python/td_common/model",
        "//python/td_common/repository",
        "@pypi//pymongo",
        "@pypi//pytest",
        "@pypi//testcontainers",
    ],
)