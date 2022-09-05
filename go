#!/bin/bash

set -e

# >>> conda initialize >>>
# !! Contents within this block are managed by 'conda init' !!
__conda_setup="$('/Users/kyunghoon/miniconda3/bin/conda' 'shell.zsh' 'hook' 2> /dev/null)"
if [ $? -eq 0 ]; then
    eval "$__conda_setup"
else
    if [ -f "/Users/kyunghoon/miniconda3/etc/profile.d/conda.sh" ]; then
        . "/Users/kyunghoon/miniconda3/etc/profile.d/conda.sh"
    else
        export PATH="/Users/kyunghoon/miniconda3/bin:$PATH"
    fi
fi
unset __conda_setup
# <<< conda initialize <<<

conda activate ldm

python scripts/dream.py --full_precision
