{{/*
Common labels
*/}}
{{- define "promtail.labels" -}}
helm.sh/chart: {{ include "promtail.chart" . }}
giantswarm.io/service-type: "managed"
{{ include "promtail.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
application.giantswarm.io/team: {{ index .Chart.Annotations "application.giantswarm.io/team" | default "atlas" | quote }}
{{- end }}
