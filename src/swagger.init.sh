#!/bin/bash

command -v swagger && (
    swagger init spec \
    --title "flight monitor application" \
    --version 1.0.0 \
    --scheme http \
    --consumes application/io.goswagger.examples.todo-list.v1+json \
    --produces application/io.goswagger.examples.todo-list.v1+json
)