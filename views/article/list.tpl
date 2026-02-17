{{template "layout/base.tpl" .}}
{{define "content"}}
<div class="mb-6">
  <h1 class="text-2xl md:text-3xl font-bold text-gray-800">文章中心</h1>
  <p class="text-gray-500 mt-2">所有已发布文章</p>
</div>
<div class="flex gap-2 flex-wrap mb-8">
  <a href="/article" class="px-3 py-1.5 rounded-full text-sm {{if eq .Cate ""}}bg-indigo-600 text-white{{else}}bg-gray-100 text-gray-700{{end}}">全部</a>
  <a href="/article?cate=技术" class="px-3 py-1.5 rounded-full text-sm {{if eq .Cate "技术"}}bg-indigo-600 text-white{{else}}bg-gray-100 text-gray-700{{end}}">技术</a>
  <a href="/article?cate=生活" class="px-3 py-1.5 rounded-full text-sm {{if eq .Cate "生活"}}bg-indigo-600 text-white{{else}}bg-gray-100 text-gray-700{{end}}">生活</a>
  <a href="/article?cate=游戏" class="px-3 py-1.5 rounded-full text-sm {{if eq .Cate "游戏"}}bg-indigo-600 text-white{{else}}bg-gray-100 text-gray-700{{end}}">游戏</a>
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
{{end}}
