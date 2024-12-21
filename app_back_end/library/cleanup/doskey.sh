#!/bin/bash

unalias -a

unset -f $(declare -F | awk '{print $3}')

exec bash
