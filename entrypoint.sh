#!/bin/bash
# Dynamically configure telegraf to run based on the need configuration.
# This is a bit dirty....but it works.
# TODO: support for getting the telegraf configuration from Git?

TELEGRAF_CONFIGURE="${TELEGRAF_CONFIGURE:-1}"
TELEGRAF_CONFIG_PATH="${TELEGRAF_CONFIG_PATH:-/metrics/telegraf.conf}"
if [ "${TELEGRAF_CONFIGURE}" == "yes" ]; then
    cat <<EOC
=========================================================================
Sidecar container: $(/bin/telegraf --version)
EOC
    mkdir -p $(dirname ${TELEGRAF_CONFIG_PATH})
    cp /etc/telegraf/recurly.conf ${TELEGRAF_CONFIG_PATH}

    cat >> ${TELEGRAF_CONFIG_PATH} <<EOM
# Graphite configuration is dynamically writen.
[[outputs.graphite]]
    servers = ["${GRAPHITE_SRV1}", "${GRAPHITE_SRV2}"]
    prefix = "${LOGGING_NAMESPACE}"

# The following configuration was run-time injected from
# the environment variable TELEGRAF_CONFIG_PATH
${TELEGRAF_PLUGIN_CONFIG}

EOM
fi

if [ $# -gt 0 ]; then
    exec ${@}
fi

exec /bin/telegraf
