{{template "layout/base.tpl" .}}
{{define "content"}}
<div class="mb-6">
  <h1 class="text-2xl font-bold text-gray-800">{{.Tool.Title}}</h1>
</div>
<div class="bg-white p-4 rounded-2xl shadow-sm">
  <iframe src="/static/uploads/{{.Tool.Folder}}/index.html" width="100%" height="600" frameborder="0"></iframe>
</div>
{{end}}
