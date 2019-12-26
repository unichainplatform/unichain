# Copyright 2018 The UniChain Team Authors
# This file is part of the unichain project.
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program. If not, see <http://www.gnu.org/licenses/>.

#!/usr/bin/env bash

# Gets the git commit hash of the working dir and adds an additional hash of any tracked modified files
commit=$(git describe --tags)
dirty=$(git ls-files -m)
branch=$(git branch | grep '*' | cut -d ' ' -f 2)

commit="$commit+branch.$branch"
if [[ -n ${dirty} ]]; then
    commit="$commit+dirty.$(echo ${dirty} | git hash-object --stdin | head -c8)"
fi
echo "$commit"

