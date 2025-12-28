from setuptools import setup, find_packages
import pathlib

here = pathlib.Path(__file__).parent.resolve()

# Read the README file
long_description = (here / "README.md").read_text(encoding="utf-8")

setup(
    name="smerch",
    version="0.0.2",
    description="A high-performance no-GIL Python package manager",
    long_description=long_description,
    long_description_content_type="text/markdown",
    author="joegasewicz",
    author_email="contact@josef.digital",
    python_requires=">=3.14",
    packages=find_packages(),
    install_requires=[
        # add runtime dependencies here if needed
    ],
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
)