{{template "layout/base.tpl" .}}
{{define "content"}}
<div class="flex justify-between items-center mb-6">
  <h1 class="text-2xl font-bold">文章管理</h1>
  <a href="/admin/article/add" class="bg-indigo-600 text-white px-4 py-2 rounded">+ 新增文章</a>
</div>
<div class="bg-white rounded shadow overflow-hidden">
<table class="w-full">
<tr class="bg-gray-100">
  <th class="p-3 text-left">ID</th>
  <th class="p-3 text-left">标题</th>
  <th class="p-3 text-left">分类</th>
  <th class="p-3 text-left">状态</th>
  <th class="p-3 text-left">时间</th>
  <th class="p-3 text-left">操作</th>
</tr>
{{range .Articles}}
<tr class="border-t">
  <td class="p-3">{{.Id}}</td>
  <td class="p-3">{{.Title}}</td>
  <td class="p-3">{{.Category}}</td>
  <td class="p-3">{{if eq .Status 2}}<span class="text-green-600">已发布</span>{{else}}草稿{{end}}</td>
  <td class="p-3 text-sm text-gray-500">{{.CreateTime.Format "2006-01-02"}}</td>
  <td class="p-3 space-x-2">
    <a href="/admin/article/edit/{{.Id}}" class="text-blue-600">编辑</a>
    <a href="/admin/article/del/{{.Id}}" class="text-red-600" onclick="return confirm('确定删除？')">删除</a>
  </td>
</tr>
{{end}}
</table>
</div>
{{end}}
