#!/bin/bash
#
# Copyright 2023 Ant Group Co., Ltd.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

set -e
usage="$(basename "$0") THIRDTYPE"
THIRDTYPE=$1
PARTY_ID=$2
OTHER_PARTY_ID=$3
OTHER_PARTY_IP=$4

if [[ ${THIRDTYPE} == "" ]]; then
    echo "missing argument: $usage"
    exit 1
fi

ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)

if [[ ${THIRDTYPE} == "fate" ]]; then
    su - app -s ${ROOT}/scripts/deploy_fate_cluster.sh ${PARTY_ID} ${OTHER_PARTY_ID} ${OTHER_PARTY_IP}
fi

while true
do
    sleep 3600
done