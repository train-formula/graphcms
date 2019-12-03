#!/usr/bin/env bash


# Dependencies: This script requires AWS CLI 1.5.3+, jq, and ruby 1.9+
# This script ONLY WORKS WITH K/V SECRETS!
# THIS IS MEANT TO BE CALLED WITH THE LOGIC CONTAINED IN A DEPLOY BASH SCRIPT!

# This bash script loads a secret from the AWS Secrets Manager
# It parses the output with jq
# Then it parses the JSON string stored with jq again
# Then it pipes that output into Ruby
# Ruby takes it, loads the json, and base64 encodes the map values
# After that, it dumps it into YAML
# After that, it pipelines the YAML into sed, which drops the first line
# We need to do this as Ruby appends --- to the output, but to add it to an existing YAML file we don't want to do this

SECRET_NAMES=("graphcms-staging")
SECRET_FILE=`pwd`/secrets.yaml

# Specify some dependencies with some defaults that imagine we used homebrew to install things. bad i know :(
system_jq=$(which jq)
JQ=${system_jq:-/usr/local/bin/jq}
system_aws=$(which aws)
AWS=${system_aws:-/usr/local/bin/aws}
system_ruby=$(which ruby)
RUBY=${system_ruby:-/usr/local/bin/ruby}

# clear out the secret file, on the off chance the secrets file persists from another load, because we use >> (append)
: > ${SECRET_FILE} | true

for i in ${SECRET_NAMES[*]};
do
    ${AWS} secretsmanager get-secret-value --secret-id ${i} | \
    ${JQ} -rc .SecretString | \
    ${RUBY} -ryaml -rjson -rbase64 -e 'puts YAML.dump(Hash[JSON.parse(STDIN.gets).map{|k,str| [k,Base64.strict_encode64(str)] }])' | \
    sed -n '1!p' \
    >> ${SECRET_FILE}
done
