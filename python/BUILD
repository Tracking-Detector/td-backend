load("@rules_python//python:pip.bzl", "compile_pip_requirements")

compile_pip_requirements(
    name = "requirements",
    src = "requirements.in",
    extra_args = [
        "--quiet",
    ],
    requirements_txt = "requirements_lock.txt",
    requirements_windows = "requirements_windows.txt",
    tags = ["no_test"],
)
