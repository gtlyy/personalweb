{{template "layout/base.tpl" .}}
{{define "content"}}
<div class="mb-6">
  <h1 class="text-2xl md:text-3xl font-bold text-gray-800">在线游戏</h1>
  <p class="text-gray-500 mt-2">无需下载，直接游玩</p>
</div>
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
  {{range .Games}}
  <div class="bg-white rounded-2xl shadow-sm hover:shadow-md hover:-translate-y-1 transition-all">
    <div class="p-6">
      <span class="inline-block px-2 py-1 text-xs bg-emerald-100 text-emerald-600 rounded-full mb-3">{{.Category}}</span>
      <a href="/game/{{.Id}}">
        <h3 class="text-lg font-semibold text-gray-800 hover:text-emerald-600">{{.Title}}</h3>
      </a>
      <p class="text-gray-400 text-sm mt-3">{{.CreateTime.Format "2006-01-02"}}</p>
    </div>
  </div>
  {{end}}
</div>
{{end}}
