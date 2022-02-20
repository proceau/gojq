gojq() {
  local yaml="--yaml-input"
  local arg
  for arg in "$@"; do
    if [[ "$arg" = "--json-input" ]]; then
      unset yaml
    fi
  done
  command gojq ${yaml:+"$yaml"} "${@:#--json-input}"
}
