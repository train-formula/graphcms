
{{- define "deploymentName" -}}
{{ printf "%s" .Chart.Name }}
{{- end -}}

{{- define "chartName" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
	Creates a sane progressDeadlineSeconds default based on our minReadySeconds value

	progressDeadlineSeconds tells the a kubernetes rolling-update how many seconds of inactivity to consider a failure
*/}}

{{- define "progressDeadlineSeconds" -}}
{{- add .Values.minReadySeconds 5 -}}
{{- end -}}


{{- define "commonLabels" -}}
app: {{ template "deploymentName" . }}
chart: {{ template "chartName" . }}
heritage: {{ .Release.Service | quote }}
release: {{ .Release.Name | quote }}
environment: {{ .Values.environment | quote }}
component: "main"
{{- end -}}