#!/bin/sh

set -eou pipefail

# This script is run by the top-level Makefile
# It's intended to be run from the root as ./scripts/gen_arm.sh
#
# It generates an ARM struct for the template at /arm/aci_template.json
#
# requires gojson to be installed. See https://github.com/ChimeraCoder/gojson
# for more information

ARM_INPUT_FILE=${PWD}/arm/aci_template.json
ARM_OUTPUT_FILE=${PWD}/aci/aci_arm.go
ARM_FILE=$(cat ${ARM_INPUT_FILE})

echo "Generating struct ${ARM_OUTPUT_FILE} from JSON ${ARM_INPUT_FILE}"
gojson \
-name=aciARMTpl \
-input ${ARM_INPUT_FILE} \
-pkg aci \
-subStruct=true \
> ${ARM_OUTPUT_FILE}
