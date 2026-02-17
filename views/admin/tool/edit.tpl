{{template "layout/base.tpl" .}}
{{define "content"}}
<h1 class="text-2xl font-bold mb-6">编辑工具</h1>
<div class="bg-white p-6 rounded shadow">
<form method="post">
  <div class="mb-4">
    <label class="block mb-1 font-medium">名称</label>
    <input name="title" value="{{.Tool.Title}}" class="w-full border px-3 py-2 rounded" required>
  </div>
  <div class="mb-4">
    <label class="block mb-1 font-medium">分类</label>
    <input name="category" value="{{.Tool.Category}}" class="w-full border px-3 py-2 rounded" required>
  </div>
  <div class="mb-4">
    <label class="block mb-1 font-medium">状态</label>
    <label class="mr-4"><input type="radio" name="status" value="1" {{if eq .Tool.Status 1}}checked{{end}}>草稿</label>
    <label><input type="radio" name="status" value="2" {{if eq .Tool.Status 2}}checked{{end}}>发布</label>
  </div>
  <button class="bg-blue-600 text-white px-4 py-2 rounded">保存</button>
</form>
</div>
{{end}}
