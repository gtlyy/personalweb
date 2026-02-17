{{template "layout/base.tpl" .}}
{{define "content"}}
<div class="py-16 md:py-24 text-center">
  <h1 class="text-4xl md:text-5xl font-bold text-gray-800 mb-4">个人主页</h1>
  <p class="text-gray-500 text-lg max-w-xl mx-auto">文章、在线小游戏、实用工具，简单好用</p>
  <div class="mt-8 flex justify-center gap-4 flex-wrap">
    <a href="/article" class="px-6 py-3 bg-indigo-600 text-white rounded-xl hover:bg-indigo-700 transition">浏览文章</a>
    <a href="/games" class="px-6 py-3 bg-emerald-600 text-white rounded-xl hover:bg-emerald-700 transition">在线游戏</a>
    <a href="/tools" class="px-6 py-3 bg-blue-600 text-white rounded-xl hover:bg-blue-700 transition">实用工具</a>
  </div>
</div>

<div class="mt-10">
  <div class="flex justify-between items-center mb-6">
    <h2 class="text-2xl font-bold text-gray-800">最新文章</h2>
    <a href="/article" class="text-indigo-600 hover:underline">查看全部</a>
  </div>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
    {{range .Articles}}
    <div class="bg-white rounded-2xl shadow-sm hover:shadow-md hover:-translate-y-1 transition-all">
      <div class="p-6">
        <span class="inline-block px-2 py-1 text-xs bg-indigo-100 text-indigo-600 rounded-full mb-3">{{.Category}}</span>
        <a href="/article/{{.Id}}">
          <h3 class="text-lg font-semibold text-gray-800 hover:text-indigo-600 line-clamp-2">{{.Title}}</h3>
        </a>
        <p class="text-gray-400 text-sm mt-3">{{.CreateTime.Format "2006-01-02"}}</p>
      </div>
    </div>
    {{end}}
  </div>
</div>
{{end}}
